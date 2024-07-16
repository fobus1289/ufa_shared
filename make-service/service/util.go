package service

import (
	"github.com/iancoleman/strcase"
	"text/template"
)

func toLowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}

func toCamel(s string) string {
	return strcase.ToCamel(s)
}

func toSnake(s string) string {
	return strcase.ToSnake(s)
}

func toKebab(s string) string {
	return strcase.ToKebab(s)
}

func withSpace(s string) string {
	return strcase.ToDelimited(s, ' ')
}

func Tmp(data string) *template.Template {
	return template.Must(template.New("").Funcs(template.FuncMap{
		"toCamel":      toCamel,
		"toLowerCamel": toLowerCamel,
		"toSnake":      toSnake,
		"toKebab":      toKebab,
		"withSpace":    withSpace,
	}).Parse(data))
}
