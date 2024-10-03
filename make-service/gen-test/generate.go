package gentest

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"text/template"

	"github.com/fobus1289/ufa_shared/make-service/stuble"
	"golang.org/x/tools/imports"
)

func readJson(filePath string) (map[string]any, error) {
	data, err := os.ReadFile(filePath + "/swagger.json")
	{
		if err != nil {
			return nil, err
		}
	}

	var swaggerJSON map[string]any
	{
		if err = json.Unmarshal(data, &swaggerJSON); err != nil {
			return nil, err
		}
	}

	return swaggerJSON, nil
}

func generateCodeByTpl(data any, tpl string) (string, error) {

	tmpl := template.New("test")

	tmpl.Funcs(template.FuncMap{
		"title":               titleCase,
		"cut":                 cutWord,
		"upper":               upper,
		"extractDtoName":      extractDtoName,
		"swaggerTypeToGoType": swaggerTypeToGoType,
		"toPascalCase":        ToPascalCase,
		"makeFnName":          makeFnName,
		"getAPIPath":          getAPIPath,
		"between":             between,
	})

	tmpl, err := tmpl.Parse(tpl)
	{
		if err != nil {
			return "", err
		}
	}

	var buff bytes.Buffer
	{
		if err = tmpl.Execute(&buff, data); err != nil {
			return "", err
		}
	}

	//fmt.Println(buff.String())
	dataimport, err := imports.Process("", buff.Bytes(), &imports.Options{
		Fragment:   false,
		FormatOnly: false,
		Comments:   true,
	})
	{
		if err != nil {
			return "", err
		}
	}

	//	return buff.String(), nil
	return string(dataimport), nil
}

// Write generated Go code to file
func writeToFile(content string, path string) error {
	f, err := os.Create(path + "/api_test.go")
	{
		if err != nil {
			return err
		}
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}

func GenerateTest(swaggFilePath, testPath string) {

	jsData, err := readJson(swaggFilePath)
	{
		if err != nil {
			log.Fatal(err)
		}
	}

	content, err := generateCodeByTpl(jsData, stuble.GoTest)
	{
		if err != nil {
			log.Fatal(err)
		}
	}

	err = writeToFile(content, testPath)
	{
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Generated test code successfully!")
}
