package internal

import (
	"errors"
	"fmt"
	"github.com/fobus1289/ufa_shared/make-service/service"
	"github.com/fobus1289/ufa_shared/make-service/stuble"
	"log"
	"path"
)

func NewService(serviceName string) {
	serviceDir := serviceName + "_service"

	if service.Exists(serviceDir) {
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

	if err := service.CreateFolders(dirs); err != nil {
		log.Fatalln(err)
	}

	files := map[string]string{
		path.Join(serviceName+"_service", "cmd/main.go"):                             stuble.Cmd,
		path.Join(serviceName+"_service", serviceName, "dto", serviceName+".go"):     stuble.Dto,
		path.Join(serviceName+"_service", serviceName, "model", serviceName+".go"):   stuble.Model,
		path.Join(serviceName+"_service", serviceName, "service", serviceName+".go"): stuble.Service,
		path.Join(serviceName+"_service", serviceName, "handler", serviceName+".go"): stuble.Handler,
		path.Join(serviceName+"_service", "transport", "service", "http.go"):         stuble.Http,
		path.Join(serviceName+"_service", ".gitignore"):                              stuble.Gitignore,
		path.Join(serviceName+"_service", ".env"):                                    stuble.Env,
		path.Join(serviceName+"_service", "example.env"):                             stuble.Env,
		path.Join(serviceName+"_service", "README.md"):                               stuble.README,
	}

	if err := service.CreateFiles(serviceName, files); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s created successfully\n", serviceDir)

	if err := service.InitProject(serviceName); err != nil {
		log.Fatalln(err)
	}

	log.Println("done")
}
