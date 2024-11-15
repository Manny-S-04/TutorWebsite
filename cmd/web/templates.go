package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/url"
	"path/filepath"
	"website/pkg/models"
	getEmbedded "website/ui"
)


type templateData struct{
    FormData url.Values
    FormErrors map[string]string
    Reviews []*models.Review
}

func newTemplateCache() (map[string]*template.Template, error) {
    embeddedHTML := getEmbedded.GetEmbeddedHTML()

    fmt.Println(embeddedHTML)

	cache := map[string]*template.Template{}

    pages, err := fs.Glob(embeddedHTML, "html/*.page.tmpl")
	if err != nil {
		return nil, err
	}
    
    fmt.Println(pages)

	for _, page := range pages {
		name := filepath.Base(page)
        fmt.Println(name)

        ts, err := template.New(name).ParseFS(embeddedHTML, "html/" +name)
        if err != nil{
            return nil, err
        }
        ts, err = template.New(name).ParseFS(embeddedHTML, "html/*.layout.tmpl")
        if err != nil{
            return nil, err
        }
        ts, err = template.New(name).ParseFS(embeddedHTML, "html/*.partial.tmpl")
        if err != nil{
            return nil, err
        }


		cache[name] = ts
	}
	return cache, nil
}
