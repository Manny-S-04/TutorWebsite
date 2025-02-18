package embed

import (
	"embed"
	"io/fs"
)

var (
	//go:embed all:static
	static      embed.FS
)

func GetStatic() embed.FS{
    return static
}

func GetStaticDirFS() (fs.FS, error) {
    return fs.Sub(static, "static")
}

