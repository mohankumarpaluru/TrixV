package app

import (
	"encoding/binary"
	"fmt"

	log "github.com/sirupsen/logrus"

	"go.mills.io/bitcask/v2"
)

// BitcaskStore ...
type BitcaskStore struct {
	db bitcask.DB
}

// NewBitcaskStore ...
func NewBitcaskStore(path string, options ...bitcask.Option) (Store, error) {
	db, err := bitcask.Open(path, options...)
	if err != nil {
		return nil, err
	}
	return &BitcaskStore{db: db}, nil
}

// Migrate ...
func (s *BitcaskStore) Migrate(collection, id string) error {
	if s.db.Has([]byte(fmt.Sprintf("/views/%s/%s", collection, id))) {
		oldViews, err := s.GetViews_(collection, id)
		if err != nil {
			err := fmt.Errorf("error getting old views for %s %s: %w", collection, id, err)
			return err
		}

		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(oldViews))
		err = s.db.Put([]byte(fmt.Sprintf("/views/%s/%s", collection, id)), buf)
		if err != nil {
			err := fmt.Errorf("error storing new views for %s: %w", id, err)
			return err
		}

		if err := s.db.Delete([]byte(fmt.Sprintf("/views/%s/%s", collection, id))); err != nil {
			err := fmt.Errorf("error deleting old views for %s %s: %w", collection, id, err)
			return err
		}
	}
	return nil
}

// Close ...
func (s *BitcaskStore) Close() error {
	return s.db.Close()
}

// GetViews_ ...
func (s *BitcaskStore) GetViews_(collection, id string) (int64, error) {
	var views uint64
	rawViews, err := s.db.Get([]byte(fmt.Sprintf("/views/%s/%s", collection, id)))
	if err != nil {
		if err != bitcask.ErrKeyNotFound {
			err := fmt.Errorf("error getting views for %s %s: %w", collection, id, err)
			log.Error(err)
			return 0, err
		}
	} else {
		views = binary.BigEndian.Uint64(rawViews)
	}

	return int64(views), nil
}

// IncView_ ...
func (s *BitcaskStore) IncView_(collection, id string) error {
	views, err := s.GetViews_(collection, id)
	if err != nil {
		err := fmt.Errorf("error getting existing views for %s %s: %w", collection, id, err)
		return err
	}

	buf := make([]byte, 8)
	views++
	binary.BigEndian.PutUint64(buf, uint64(views))
	err = s.db.Put([]byte(fmt.Sprintf("/views/%s/%s", collection, id)), buf)
	if err != nil {
		err := fmt.Errorf("error storing updated views for %s %s: %w", collection, id, err)
		return err
	}

	return nil
}

// GetViews ...
func (s *BitcaskStore) GetViews(id string) (int64, error) {
	var views uint64
	rawViews, err := s.db.Get([]byte(fmt.Sprintf("/views/%s", id)))
	if err != nil {
		if err != bitcask.ErrKeyNotFound {
			err := fmt.Errorf("error getting views for %s: %w", id, err)
			log.Error(err)
			return 0, err
		}
	} else {
		views = binary.BigEndian.Uint64(rawViews)
	}

	return int64(views), nil
}

// IncViews ...
func (s *BitcaskStore) IncViews(id string) error {
	views, err := s.GetViews(id)
	if err != nil {
		err := fmt.Errorf("error getting existing views for %s: %w", id, err)
		return err
	}

	buf := make([]byte, 8)
	views++
	binary.BigEndian.PutUint64(buf, uint64(views))
	err = s.db.Put([]byte(fmt.Sprintf("/views/%s", id)), buf)
	if err != nil {
		err := fmt.Errorf("error storing updated views for %s: %w", id, err)
		return err
	}

	return nil
}
