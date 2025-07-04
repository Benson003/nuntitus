package embeding

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed ui/static/*
var EmbeddedFiles embed.FS

// FS returns an http.FileSystem of the embedded frontend
func FS() http.FileSystem {
	subFS, err := fs.Sub(EmbeddedFiles, "ui/static")
	if err != nil {
		panic(err)
	}
	return http.FS(subFS)
}
