package templates

import (
	"embed"
)

//go:embed *.html
var templates embed.FS

// MustGetTemplate returns a template string with the given name.
func MustGetTemplate(name string) string {
	b, err := templates.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(b)
}
