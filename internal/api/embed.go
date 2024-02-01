package api

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed dist
var FrontendFS embed.FS

func FrontendContents() fs.FS {
	embeddedFiles, err := fs.Sub(FrontendFS, "dist")
	if err != nil {
		log.Fatalln(err)
	}
	return embeddedFiles
}
