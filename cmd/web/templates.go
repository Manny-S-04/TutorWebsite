package main

import (
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

    cache := map[string]*template.Template{}

    pages, err := fs.Glob(embeddedHTML, "html/*.page.tmpl")
    if err != nil {
        return nil, err
    }

    layouts, err := fs.Glob(embeddedHTML, "html/*.layout.tmpl")
    if err != nil {
        return nil, err
    }

    partials, err := fs.Glob(embeddedHTML, "html/*.partial.tmpl")
    if err != nil {
        return nil, err
    }

    for _, layout := range layouts {
        name := filepath.Base(layout)

        ts, err := template.New(name).ParseFS(embeddedHTML, "html/"+name)
        if err != nil {
            return nil, err
        }

        cache[name] = ts
    }

    for _, page := range pages {
        name := filepath.Base(page)
        ts, err := template.New(name).ParseFS(embeddedHTML, "html/"+name)
        if err != nil {
            return nil, err
        }

        for _, layout := range layouts {
            _, err := ts.ParseFS(embeddedHTML, layout)
            if err != nil {
                return nil, err
            }
        }

        for _, partial := range partials {
            _, err := ts.ParseFS(embeddedHTML, partial)
            if err != nil {
                return nil, err
            }
        }

        cache[name] = ts
    }

    return cache, nil
}
