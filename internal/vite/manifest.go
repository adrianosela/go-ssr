package vite

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
)

type ViteManifestEntry struct {
	File    string   `json:"file"`
	Name    string   `json:"name"`
	Src     string   `json:"src"`
	IsEntry bool     `json:"isEntry"`
	CSS     []string `json:"css,omitempty"`
}

type ViteManifest struct {
	bySource map[string]*ViteManifestEntry
	byName   map[string]*ViteManifestEntry
}

func LoadManifest(path string) (*ViteManifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest file at %s: %v", path, err)
	}

	var bySource map[string]*ViteManifestEntry
	if err := json.Unmarshal(data, &bySource); err != nil {
		return nil, fmt.Errorf("failed to decode manifest file as JSON: %v", err)
	}

	// map by name
	byName := make(map[string]*ViteManifestEntry, len(bySource))
	for _, v := range bySource {
		byName[v.Name] = v
	}

	return &ViteManifest{
		bySource: bySource,
		byName:   byName,
	}, nil
}

func (m *ViteManifest) LookupFileBySource(src string) (*ViteManifestEntry, bool) {
	entry, ok := m.bySource[src]
	return entry, ok
}

func (m *ViteManifest) LookupFileByName(name string) (*ViteManifestEntry, bool) {
	entry, ok := m.byName[name]
	return entry, ok
}

func (e *ViteManifestEntry) GetScriptTags() template.HTML {
	return template.HTML(fmt.Sprintf(`<script type="module" src="/static/vite/%s"></script>`, e.File))
}

func (e *ViteManifestEntry) GetStyleTags() template.HTML {
	if len(e.CSS) == 0 {
		return ""
	}
	tags := ""
	for _, cssFile := range e.CSS {
		tags += fmt.Sprintf(`<link rel="stylesheet" href="/static/vite/%s">`, cssFile)
	}
	return template.HTML(tags)
}
