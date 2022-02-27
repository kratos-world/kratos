package base

import (
	"strings"
	"unicode"
)

func SmallCamel(name string) string {
	scName := ""
	if name == "" {
		return scName
	}
	name = camelToUnderline(name)
	textList := strings.Split(name, "_")
	for idx, text := range textList {
		if text == "" {
			continue
		}
		if idx == 0 {
			scName += strings.ToLower(text)
		} else {
			scName += strings.ToUpper(string(text[0])) + strings.ToLower(text[1:])
		}
	}
	return scName
}

func BigCamel(name string) string {
	scName := ""
	if name == "" {
		return scName
	}
	name = camelToUnderline(name)
	textList := strings.Split(name, "_")
	for _, text := range textList {
		if text == "" {
			continue
		}
		scName += strings.ToUpper(string(text[0])) + strings.ToLower(text[1:])
	}
	return scName
}

func camelToUnderline(name string) string {
	var texts []string
	for idx, n := range name {
		if unicode.IsUpper(n) {
			if idx > 0 {
				texts = append(texts, "_")
			}
			texts = append(texts, strings.ToLower(string(n)))
		} else {
			texts = append(texts, string(n))
		}
	}
	return strings.Join(texts, "")
}
