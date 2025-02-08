package assets

import (
  resource "github.com/quasilyte/ebitengine-resource"
)

const (
  FontNone resource.FontID = iota
  FontNormal
  FontBig
)

func registerFontResources(loader *resource.Loader) {
  fontResources := map[resource.FontID]resource.FontInfo{
    FontNormal: {Path: "fonts/dizzy_font.ttf", Size: 14},
    FontBig:    {Path: "fonts/dizzy_font.ttf", Size: 18},
  }

  for id, res := range fontResources {
    loader.FontRegistry.Set(id, res)
    loader.LoadFont(id)
  }
}
