package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"strings"
	"unicode"

	"github.com/fobus1289/ufa_shared/make-service/stuble"
)

//PkgServiceDto
//PkgServiceModel

func main() {
	// sdir, _ := stuble.Templates.ReadDir(".")

	// for _, dir := range sdir {
	// 	if dir.IsDir() {
	// 		log.Println(dir.Name())
	// 	}
	// }

	// return

	if len(os.Args) <= 1 {
		log.Fatalln("not enough arguments")
	}

	switch os.Args[1] {
	case "--new":
		projectName := promptInput("Enter project name: ")
		modPath := promptInput("Enter project mod path: ")
		names, _ := stuble.ReadDirs(".")

		fmt.Println(strings.Join(names, "\n"))

		architecture := promptInput("Enter project architecture: ")

		if !slices.Contains(names, architecture) {
			log.Fatalf("unknown architecture => %s", architecture)
		}

		d, _ := os.Getwd()

		stuble.WalkPrintFS(stuble.Templates, architecture, projectName, d, 0)
		// structures, err := stuble.WalkEmbedFS(stuble.Templates, architecture)
		// {
		// 	if err != nil {
		// 		log.Fatalf("unknown architecture => %s", architecture)
		// 	}

		// 	fmt.Println(strings.Join(structures, "\n"))
		// }

		result := fmt.Sprintf("project name %s\nproject mod %s\nproject architecture %s\n", projectName, modPath, architecture)

		fmt.Println(result)
	}

	// switch os.Args[1] {
	// case "--new":
	// 	projectName := promptInput("Enter project name: ")
	// 	modPath := promptInput("Enter project mod path: ")
	// 	internal.NewService(projectName, modPath)
	// case "--add":
	// 	projectName := promptInput("Enter project name: ")
	// 	modPath := promptInput("Enter project mod path: ")
	// 	internal.AddService(projectName, modPath)
	// default:
	// 	log.Fatalln(errors.New("unknown flag"))
	// }
}

func promptInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	{
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		return strings.TrimSpace(input)
	}
}

type Container map[string]any

type Service struct {
	Name      string
	Mod       string
	Container Container
}

func (s *Service) UpperName() {

}

func (s *Service) LowwerName() {

}

func (s *Service) ProjectName(n string) string {
	return s.Name + strings.TrimSpace(n)
}

func (s *Service) ModuleName(n ...string) string {
	pkg := append([]string{s.Mod}, n...)
	log.Println(pkg)
	return path.Join(pkg...)
}

func (s *Service) Set(key string, value any) any {
	if s.Container == nil {
		s.Container = Container{}
	}

	s.Container[key] = value

	return ""
}

func (s *Service) Get(key string) any {
	return s.Container[key]
}

// UcFirst делает первую букву строки заглавной
func UcFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// LcFirst делает первую букву строки строчной
func LcFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}
