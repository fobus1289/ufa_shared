package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/fobus1289/ufa_shared/make-service/stuble"
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

	if len(os.Args) < 2 {
		log.Fatalln(errors.New("write service name"))
	}

	serviceName := os.Args[1]

	serviceDir := serviceName + "_service"

	if serviceExists(serviceDir) {
		log.Fatalln(errors.New("not empty service"))
	}

	dirs := []string{
		"cmd",
		"dto",
		"models",
		"service",
		"transport",
		"transport/http",
		"transport/service",
	}

	for _, dir := range dirs {
		buildDir := path.Join(serviceDir, dir)
		if err := os.MkdirAll(buildDir, 0750); err != nil {
			log.Fatalln(err)
		}
	}

	services := map[string]string{
		"cmd/main.go":                            stuble.Cmd,
		"dto/dto.go":                             stuble.Dto,
		"models/model.go":                        stuble.Model,
		"service/service.go":                     stuble.Service,
		"transport/http/transport_http.go":       stuble.TransportHttp,
		"transport/service/transport_service.go": stuble.TransportService,
		".gitignore":                             stuble.Gitignore,
		".env":                                   stuble.Env,
		"README.md":                              stuble.README,
	}

	for k, v := range services {
		buildFile := path.Join(serviceDir, k)
		if file, err := os.Create(buildFile); err != nil {
			log.Fatalln(err)
		} else {
			defer file.Close()

			m := map[string]string{
				"ServiceName": serviceName,
			}

			var buffer bytes.Buffer

			if err := tmp(v).Execute(&buffer, m); err != nil {
				log.Fatalln(err)
			}

			if _, err := file.Write(buffer.Bytes()); err != nil {
				log.Fatalln(err)
			}
		}
	}

	fmt.Printf("write command\n")
	fmt.Printf("%s service created successfully\n", serviceDir)
	fmt.Printf("cd %s\ngo mod init %s\ngo mod tidy\n", serviceDir, serviceDir)

	//TODO:auto install
	// if _, err := exec.LookPath("go"); err != nil {
	// 	fmt.Printf("install golang\n")
	// 	fmt.Printf("%s_service service created successfully\n", serviceName)
	// 	fmt.Printf("cd %s_service\ngo mod init %s_service\ngo mod tidy\n", serviceName, serviceName)
	// } else {
	// 	cmd := exec.Command("go", "mod", "init", fmt.Sprintf("%s_service", serviceName))

	// 	cmd.Stderr = os.Stderr

	// 	cmd.Path = "./" + serviceName

	// 	cmd.Start()
	// 	cmd.Wait()
	// }

}

func serviceExists(serviceName string) bool {

	info, err := os.Stat(serviceName)

	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir() && !os.IsNotExist(err)
}
