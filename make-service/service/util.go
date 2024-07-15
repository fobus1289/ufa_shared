package service

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"text/template"
)

func lcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	uc := ucFirst(s)
	return strings.ToLower(uc[:1]) + uc[1:]
}

func ucFirst(s string) string {
	parts := strings.Split(s, "_")
	caser := cases.Title(language.Tag{}, cases.NoLower)
	for i, part := range parts {
		parts[i] = caser.String(part)
	}
	return strings.Join(parts, "")
}

func Tmp(data string) *template.Template {
	return template.Must(template.New("").Funcs(template.FuncMap{
		"ucFirst": ucFirst,
		"lcFirst": lcFirst,
	}).Parse(data))
}
