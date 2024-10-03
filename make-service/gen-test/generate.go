package gentest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

func writeToFile(content string, path string) error {
	filePath := filepath.Join(path, "api_test.go")

	file, err := os.Open(filePath)
	if err == nil {
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			return fmt.Errorf("failed to get file info: %w", err)
		}

		if fileInfo.Size() > 0 {
			return fmt.Errorf("test file already exists and is not empty")
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("failed to open file %s: %w", filePath, err)
	}

	// File is either new or empty, create or truncate the file
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer f.Close()

	// Write the new content to the file
	_, err = f.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write content to file %s: %w", filePath, err)
	}

	return nil
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
