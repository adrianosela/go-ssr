package vite

import (
	"encoding/json"
	"fmt"
	"html/template"
)

const (
	funcToJSON         = "toJSON"
	funcViteScriptTags = "viteScriptTags"
	funcViteStyleTags  = "viteStyleTags"
)

func GetAllHelpers(manifest *ViteManifest) template.FuncMap {
	return map[string]any{
		funcToJSON:         getFuncToJSON(),
		funcViteScriptTags: getFuncViteScriptTags(manifest),
		funcViteStyleTags:  getFuncViteStyleTags(manifest),
	}
}

func getFuncToJSON() func(any) template.JS {
	return func(v any) template.JS {
		b, _ := json.Marshal(v)
		return template.JS(b)
	}
}

func getFuncViteScriptTags(manifest *ViteManifest) func(string) template.HTML {
	return func(name string) template.HTML {
		entry, ok := manifest.LookupFileByName(name)
		if !ok {
			return getConsoleLogTag(fmt.Errorf("no vite manifest entry for %s", name))
		}
		return entry.GetScriptTags()
	}
}

func getFuncViteStyleTags(manifest *ViteManifest) func(string) template.HTML {
	return func(name string) template.HTML {
		entry, ok := manifest.LookupFileByName(name)
		if !ok {
			return getConsoleLogTag(fmt.Errorf("no vite manifest entry for %s", name))
		}
		return entry.GetStyleTags()
	}
}

func getConsoleLogTag(err error) template.HTML {
	return template.HTML(fmt.Sprintf(`<script>console.log("%s")</script>`, err.Error()))
}
