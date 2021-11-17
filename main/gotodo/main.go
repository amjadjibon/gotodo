package main

import (
	_ "embed"

	"github.com/mkawserm/abesh/cmd"
	_ "github.com/mkawserm/httpserver2/capability/httpserver2"

	_ "github.com/amjadjibon/gotodo/capability/createalbum"
	_ "github.com/amjadjibon/gotodo/capability/getalbumbyid"
	_ "github.com/amjadjibon/gotodo/capability/getalbums"
	_ "github.com/amjadjibon/gotodo/capability/updatealbum"
)

//go:embed manifest.yaml
var manifestBytes []byte

func main() {
	cmd.ManifestBytes = manifestBytes
	cmd.Execute()
}
