package embed

import (
	"embed"
	"fmt"
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

func ListEmbeddedFiles() {
	err := fs.WalkDir(static, "static", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		if d.IsDir() {
			fmt.Println("[DIR] " + path)
		} else {
			fmt.Println("[FILE] " + path)
		}

		return nil
	})

	if err != nil {
        panic(err)
	}
}
