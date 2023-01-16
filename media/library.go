package media

import (
	"errors"
	"fmt"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

// Library manages importing and retrieving video data.
type Library struct {
	mu     sync.RWMutex
	Paths  map[string]*Path
	Videos map[string]*Video
}

// NewLibrary returns new instance of Library.
func NewLibrary() *Library {
	lib := &Library{
		Paths:  make(map[string]*Path),
		Videos: make(map[string]*Video),
	}
	return lib
}

// AddPath adds a media path to the library.
func (lib *Library) AddPath(p *Path) error {
	lib.mu.Lock()
	defer lib.mu.Unlock()
	// make sure new path doesn't collide with existing ones
	for _, p2 := range lib.Paths {
		if p.Path == p2.Path {
			return errors.New(fmt.Sprintf("media: duplicate (normalized) library path '%s'", p.Path))
		}
		if p.Prefix == p2.Prefix {
			return errors.New(fmt.Sprintf("media: duplicate library prefix '%s'", p.Prefix))
		}
	}
	if _, err := os.Stat(p.Path) ; err != nil && os.IsNotExist(err) {
		log.Warn(fmt.Sprintf("media: library path '%s' does not exist. Creating it now.", p.Path))
		if err := os.MkdirAll(p.Path, 0o755); err != nil {
			return fmt.Errorf("error creating library path %s: %w", p.Path, err)
		}
	}
	lib.Paths[p.Path] = p
	return nil
}

// Import adds all valid videos from a given path.
func (lib *Library) Import(p *Path) error {
	files, err := ioutil.ReadDir(p.Path)
	if err != nil {
		return err
	}
	for _, info := range files {
		if strings.ContainsAny(info.Name(), "#") {
			// ignore resized videos e.g: #240p.mp4
			continue
		}
		err = lib.Add(path.Join(p.Path, info.Name()))
		if err != nil {
			// Ignore files that can't be parsed
			continue
		}
	}
	return nil
}

// Add adds a single video from a given file path.
func (lib *Library) Add(fp string) error {
	lib.mu.Lock()
	defer lib.mu.Unlock()
	fp = filepath.ToSlash(fp)
	d := path.Dir(fp)
	p, ok := lib.Paths[d]
	if !ok {
		return errors.New("media: path not found")
	}
	n := path.Base(fp)
	v, err := ParseVideo(p, n)
	if err != nil {
		return err
	}
	lib.Videos[v.ID] = v
	log.Debug("Added:", v.Path)
	return nil
}

// Remove removes a single video from a given file path.
func (lib *Library) Remove(fp string) {
	lib.mu.Lock()
	defer lib.mu.Unlock()
	fp = filepath.ToSlash(fp)
	d := path.Dir(fp)
	p, ok := lib.Paths[d]
	if !ok {
		return
	}
	n := path.Base(fp)
	// ID is name without extension
	idx := strings.LastIndex(n, ".")
	if idx == -1 {
		idx = len(n)
	}
	id := n[:idx]
	if len(p.Prefix) > 0 {
		id = path.Join(p.Prefix, id)
	}
	v, ok := lib.Videos[id]
	if ok {
		delete(lib.Videos, id)
		log.Debug("Removed:", v.Path)
	}
}

// Playlist returns a sorted Playlist of all videos.
func (lib *Library) Playlist() Playlist {
	lib.mu.RLock()
	defer lib.mu.RUnlock()
	pl := make(Playlist, len(lib.Videos))
	i := 0
	for _, v := range lib.Videos {
		pl[i] = v
		i++
	}
	By(SortByTimestamp).Sort(pl)
	return pl
}
