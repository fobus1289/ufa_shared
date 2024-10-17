package genfrontend

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/fobus1289/ufa_shared/make-service/stuble"
)

func generateCode() (string, error) {

	var (
		buff bytes.Buffer
	)

	tmpl := template.New("test")
	{
		tmpl.Funcs(template.FuncMap{
			"getTitle": getTitle,
		})

		tmpl, err := tmpl.Parse(stuble.GoTemplateAutoFront)
		{
			if err != nil {
				return "", err
			}
		}

		{
			if err = tmpl.Execute(&buff, nil); err != nil {
				return "", err
			}
		}
	}

	return buff.String(), nil
}

func writeToFile(content string, path string) error {
	filePath := filepath.Join(path, "auto_front.html")

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

	f, err := os.Create(filePath)
	{
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", filePath, err)
		}
	}
	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		return fmt.Errorf("failed to write content to file %s: %w", filePath, err)
	}

	return nil
}

func GenerateFront(swaggFilePath, path string) {

	err := ReadOpenAPISpec(swaggFilePath)
	{
		if err != nil {
			log.Fatal(err)
		}
	}

	content, err := generateCode()
	{
		if err != nil {
			log.Fatal(err)
		}
	}

	err = writeToFile(content, path)
	{
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Generated auto fromt successfully!")
}
