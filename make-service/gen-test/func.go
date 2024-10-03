package gentest

import (
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func cutWord(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 2 {
		if parts[2] == "{id}" {
			return titleCase(parts[1]) + "ById"
		}
		return titleCase(parts[1]) + titleCase(parts[2])
	}

	if len(parts) > 1 {
		return parts[1]
	}

	return path
}
func titleCase(s string) string {
	return cases.Title(language.English).String(strings.Trim(s, " "))
}

func upper(s string) string {
	return strings.ToUpper(s)
}

func extractDtoName(ref string) string {
	const prefix = "#/definitions/"

	startIndex := strings.Index(ref, prefix)
	{
		if startIndex == -1 {
			return ""
		}
	}

	startIndex += len(prefix)
	dtoName := ref[startIndex:]
	dtoName = strings.TrimSpace(dtoName)

	return dtoName
}

func swaggerTypeToGoType(swaggerType string) string {
	switch swaggerType {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "boolean":
		return "bool"
	case "array":
		return "[]any"
	case "object":
		return "map[string]any"
	default:
		return "any"
	}
}

func ToPascalCase(input string) string {
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r == ' ' || r == '_'
	})

	acronyms := map[string]string{
		"id":  "ID",
		"url": "URL",
		"api": "API",
	}

	for i := 0; i < len(words); i++ {
		word := strings.ToLower(words[i])

		if val, ok := acronyms[word]; ok {
			words[i] = val
		} else {
			words[i] = capitalizeFirstLetter(word)
		}
	}

	return strings.Join(words, "")
}

func makeFnName(path string) string {
	parts := strings.Split(path, "/")

	for i, part := range parts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			part = strings.Trim(part, "{}")
			parts[i] = "By" + strings.Title(part)
		} else {
			parts[i] = strings.Title(part)
		}
	}

	functionName := strings.Join(parts, "")

	functionName = capitalizeFirstLetter(functionName)

	return functionName
}

func capitalizeFirstLetter(s string) string {
	runes := []rune(s)
	if len(runes) > 0 && unicode.IsLower(runes[0]) {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}

func getAPIPath(basePath string) string {
	// Find the first `/` in the basePath
	slashIndex := strings.Index(basePath, "/")
	if slashIndex != -1 {
		// Return everything from the first `/` onwards
		return basePath[slashIndex:]
	}
	// Return an empty string if no `/` is found (shouldn't happen in your case)
	return ""
}

func between(value string, min, max int) bool {
	code, _ := strconv.Atoi(value)
	return code >= min && code <= max
}
