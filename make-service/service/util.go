package service

import (
	"github.com/iancoleman/strcase"
	"text/template"
)

func ToLowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}

func ToCamel(s string) string {
	return strcase.ToCamel(s)
}

func ToSnake(s string) string {
	return strcase.ToSnake(s)
}

func ToKebab(s string) string {
	return strcase.ToKebab(s)
}

func WithSpace(s string) string {
	return strcase.ToDelimited(s, ' ')
}

func Tmp(data string) *template.Template {
	return template.Must(template.New("").Funcs(template.FuncMap{
		"toCamel":      ToCamel,
		"toLowerCamel": ToLowerCamel,
		"toSnake":      ToSnake,
		"toKebab":      ToKebab,
		"withSpace":    WithSpace,
	}).Parse(data))
}
