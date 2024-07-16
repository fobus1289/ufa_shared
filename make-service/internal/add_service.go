package internal

import (
	"fmt"
	"github.com/fobus1289/ufa_shared/make-service/service"
	"github.com/fobus1289/ufa_shared/make-service/stuble"
	"log"
)

func AddService(serviceName, modPath string) {

	//TODO: check serviceName if exists
	var dirs = []string{
		serviceName,
		fmt.Sprintf("%s/dto", serviceName),
		fmt.Sprintf("%s/model", serviceName),
		fmt.Sprintf("%s/service", serviceName),
		fmt.Sprintf("%s/handler", serviceName),
	}

	if err := service.CreateFolders(dirs); err != nil {
		log.Fatalln(err)
	}

	files := map[string]string{
		fmt.Sprintf("%s/dto/%s.go", serviceName, serviceName):     stuble.Dto,
		fmt.Sprintf("%s/model/%s.go", serviceName, serviceName):   stuble.Model,
		fmt.Sprintf("%s/service/%s.go", serviceName, serviceName): stuble.Service,
		fmt.Sprintf("%s/handler/%s.go", serviceName, serviceName): stuble.Handler,
	}

	if err := service.CreateFiles(serviceName, modPath, files); err != nil {
		log.Fatalln(err)
	}

	//updateCmdMainFile()
	if err := service.UpdateMainGoFile(serviceName, modPath); err != nil {
		log.Fatalln(err)
	}

	//updateTransportHttp()
	if err := service.UpdateTransportHttpFile(serviceName, modPath); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s created successfully\n", serviceName)

	if err := service.GoModTidy("./" + serviceName); err != nil {
		log.Fatalln(err)
	}
	if err := service.RunGoImports("golang.org/x/tools/cmd/goimports@latest", "./"+serviceName); err != nil {
		log.Fatalln(err)
	}

	log.Println("done")
}
