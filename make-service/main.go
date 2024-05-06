package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/fobus1289/ufa_shared/make-service/stuble"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ucFirst(s string) string {
	return cases.Title(language.Tag{}, cases.NoLower).String(s)
}

func lcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func tmp(data string) *template.Template {
	return template.Must(template.New("").Funcs(template.FuncMap{
		"ucFirst": ucFirst,
		"lcFirst": lcFirst,
	}).Parse(data))
}

func main() {

	if len(os.Args) > 2 {
		switch os.Args[1] {
		case "--new":
			NewService(os.Args[2])
		case "--add":
			AddService(os.Args[2])
		default:
			log.Fatalln(errors.New("unknown flag"))
			//call help()
		}
	} else {
		log.Fatalln(errors.New("not enough arguments"))
	}

}

func NewService(serviceName string) {
	serviceDir := serviceName + "_service"

	if serviceExists(serviceDir) {
		log.Fatalln(errors.New("service already exists"))
	}

	dirs := []string{
		serviceName + "_service",
		path.Join(serviceName+"_service", "cmd"),
		path.Join(serviceName+"_service", serviceName),
		path.Join(serviceName+"_service", serviceName, "dto"),
		path.Join(serviceName+"_service", serviceName, "model"),
		path.Join(serviceName+"_service", serviceName, "service"),
		path.Join(serviceName+"_service", serviceName, "handler"),
		path.Join(serviceName+"_service", "transport"),
		path.Join(serviceName+"_service", "transport/service"),
	}

	if err := createFolders(dirs); err != nil {
		log.Fatalln(err)
	}

	files := map[string]string{
		path.Join(serviceName+"_service", "cmd/main.go"):                             stuble.Cmd,
		path.Join(serviceName+"_service", serviceName, "dto", serviceName+".go"):     stuble.Dto,
		path.Join(serviceName+"_service", serviceName, "model", serviceName+".go"):   stuble.Model,
		path.Join(serviceName+"_service", serviceName, "service", serviceName+".go"): stuble.Service,
		path.Join(serviceName+"_service", serviceName, "handler", serviceName+".go"): stuble.Handler,
		path.Join(serviceName+"_service", "transport", "service", "http.go"):         stuble.TransportService,
		path.Join(serviceName+"_service", ".gitignore"):                              stuble.Gitignore,
		path.Join(serviceName+"_service", ".env"):                                    stuble.Env,
		path.Join(serviceName+"_service", "example.env"):                             stuble.Env,
		path.Join(serviceName+"_service", "README.md"):                               stuble.README,
	}

	if err := createFiles(serviceName, files); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s created successfully\n", serviceDir)

	if err := initProject(serviceName); err != nil {
		log.Fatalln(err)
	}

	log.Println("done")
}

func AddService(serviceName string) {

	// check serviceName if exists
	var dirs = []string{
		serviceName,
		fmt.Sprintf("%s/dto", serviceName),
		fmt.Sprintf("%s/model", serviceName),
		fmt.Sprintf("%s/service", serviceName),
		fmt.Sprintf("%s/handler", serviceName),
	}

	if err := createFolders(dirs); err != nil {
		log.Fatalln(err)
	}

	files := map[string]string{
		fmt.Sprintf("%s/dto/%s.go", serviceName, serviceName):     stuble.Dto,
		fmt.Sprintf("%s/model/%s.go", serviceName, serviceName):   stuble.Model,
		fmt.Sprintf("%s/service/%s.go", serviceName, serviceName): stuble.Service,
		fmt.Sprintf("%s/handler/%s.go", serviceName, serviceName): stuble.Handler,
	}

	if err := createFiles(serviceName, files); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s created successfully\n", serviceName)

	//updateCmdMainFile()
	//updateTransportHttp()
	if err := goModTidy(serviceName, "./"+serviceName); err != nil {
		log.Fatalln(err)
	}

	log.Println("done")
}

func serviceExists(serviceName string) bool {

	info, err := os.Stat(serviceName)

	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir() && !os.IsNotExist(err)
}

func createFolders(dirs []string) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0750); err != nil {
			return errors.New(fmt.Sprintf("create folder error: %v", err))
		}
	}
	return nil
}

func createFiles(serviceName string, files map[string]string) error {
	for filePath, content := range files {
		file, err := os.Create(filePath)
		if err != nil {
			return errors.New(fmt.Sprintf("create file error %v", err))
		}

		m := map[string]string{
			"ServiceName": serviceName,
		}

		var buffer bytes.Buffer
		if err := tmp(content).Execute(&buffer, m); err != nil {
			file.Close()
			return errors.New(fmt.Sprintf("content copy error %v", err))
		}

		if _, err := file.Write(buffer.Bytes()); err != nil {
			file.Close()
			return errors.New(fmt.Sprintf("write content error %v", err))
		}

		file.Close()
	}
	return nil
}

func initProject(serviceName string) error {
	if err := goModInit(serviceName, "./"+serviceName+"_service"); err != nil {
		return err
	}
	if err := goModTidy(serviceName, "./"+serviceName+"_service"); err != nil {
		return err
	}
	return nil
}

func goModInit(serviceName, dir string) error {
	if _, err := exec.LookPath("go"); err != nil {
		return errors.New("go path not found, please install go")
	} else {
		cmd := exec.Command("go", "mod", "init", fmt.Sprintf("%s_service", serviceName))
		cmd.Dir = dir
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return errors.New(fmt.Sprintf("Failed to execute go mod init: %v\n", err))
		}
	}
	return nil
}

func goModTidy(serviceName, dir string) error {
	if _, err := exec.LookPath("go"); err != nil {
		return errors.New("go path not found, please install go")
	} else {
		cmd := exec.Command("go", "mod", "tidy")
		cmd.Dir = dir
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return errors.New(fmt.Sprintf("Failed to execute go mod tidy: %v\n", err))
		}
	}
	return nil
}
