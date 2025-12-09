package swaggerFiles

import (
	"embed"
	"io/fs"

	"golang.org/x/net/webdav"
)

//go:embed dist/*
var dist embed.FS

// FS holds embedded swagger ui files
var FS, _ = fs.Sub(dist, "dist")

var WebdavFS = webdav.NewMemFS()

func NewHandler() *webdav.Handler {
	return &webdav.Handler{
		FileSystem: WebdavFS,
		LockSystem: webdav.NewMemLS(),
	}
}
