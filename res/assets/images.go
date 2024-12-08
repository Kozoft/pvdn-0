package assets

import (
	resource "github.com/quasilyte/ebitengine-resource"
	_ "image/png"
)

const (
	ImageNone resource.ImageID = iota
	ImageApple
)

func registerImageResources(loader *resource.Loader) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImageApple: {Path: "images/apple-red.png"},
	}
	for id, res := range imageResources {
		loader.ImageRegistry.Set(id, res)
	}
}
