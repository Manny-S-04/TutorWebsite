package getEmbedded

import(
    "embed"
)

//go:embed html/*
var embeddedHTML embed.FS

//go:embed static/*
var embeddedStatic embed.FS

func GetEmbeddedHTML() embed.FS{
    return embeddedHTML
}

func GetEmbeddedStatic() embed.FS{
    return embeddedStatic
}

