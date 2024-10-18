package genfrontend

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/fobus1289/ufa_shared/make-service/stuble"
)

var selectedModules []string

func SelectModules() {
	var (
		modules       []string
		uniqueModules = make(map[string]bool)
		sb            strings.Builder
	)

	yellow := "\033[33m"
	reset := "\033[0m"

	modules = append(modules, "Choose module:")

	for path := range spec.Paths {
		parts := strings.Split(path, "/")
		if len(parts) > 1 && !uniqueModules[parts[1]] {
			modules = append(modules, parts[1])
			uniqueModules[parts[1]] = true
		}
	}

	sb.WriteString(fmt.Sprintf("%s%s%s\n", yellow, modules[0], reset))

	for i := 1; i < len(modules); i++ {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i, modules[i]))
	}

	fmt.Print(sb.String())

	fmt.Println("Enter integers separated by spaces:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)
	var numbers []int

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error converting input to integer:", err)
			continue
		}
		numbers = append(numbers, num)
	}

	for _, num := range numbers {
		if num > 0 && num < len(modules) {
			selectedModules = append(selectedModules, modules[num])
		} else {
			fmt.Printf("Invalid module number: %d\n", num)
		}
	}
}
func generateCode() (string, error) {

	var (
		buff bytes.Buffer
	)

	tmpl := template.New("test")
	{
		tmpl.Funcs(template.FuncMap{
			"getTitle": getTitle,
			"getPaths": getPaths,
		})

		tmpl, err := tmpl.Parse(stuble.GoTemplateAutoFront)
		{
			if err != nil {
				return "", err
			}
		}

		if err = tmpl.Execute(&buff, nil); err != nil {
			return "", err
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
			//	return fmt.Errorf("test file already exists and is not empty")
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

	SelectModules()

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

	log.Println("auto fromt Generated successfully!")
}
