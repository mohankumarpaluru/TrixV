package static

import (
	"embed"
	"net/http"
)

//go:embed *.png
//go:embed *.jpg
//go:embed *.ico
//go:embed *.js
//go:embed *.css
var files embed.FS

// MustGetFile returns the contents of a file from static as bytes.
func MustGetFile(name string) []byte {
	b, err := files.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return b
}

// GetFilesystem returns a http.FileSystem for the static files.
func GetFilesystem() http.FileSystem {
	return http.FS(files)
}
