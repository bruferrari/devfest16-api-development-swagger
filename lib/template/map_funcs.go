package template

import (
    "html/template"

    "github.com/bruferrari/devfest16-api-development-swagger/lib/context"
)

func FuncMaps() []template.FuncMap {
    return []template.FuncMap{
        map[string]interface{}{
        "Tr": context.I18n,
        }}
}