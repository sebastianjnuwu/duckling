package i18n

import (
	"embed"
	"encoding/json"
)

//go:embed locales/*.json
var files embed.FS

var Lang = "pt"

var t = map[string]map[string]string{}

func init() {
	for _, c := range []string{"en", "es", "pt"} {
		b, _ := files.ReadFile("locales/" + c + ".json")
		var m map[string]string
		_ = json.Unmarshal(b, &m)
		t[c] = m
	}
}

func Set(s string) {
	if _, ok := t[s]; ok {
		Lang = s
	}
}

func Toggle() {
	switch Lang {
	case "pt":
		Lang = "en"
	case "en":
		Lang = "es"
	default:
		Lang = "pt"
	}
}

func T(k string) string {
	if s, ok := t[Lang][k]; ok {
		return s
	}
	return k
}
