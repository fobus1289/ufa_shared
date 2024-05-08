package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/fobus1289/ufa_shared/make-service/stuble"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
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

	//TODO: check serviceName if exists
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

	//updateCmdMainFile()
	if err := updateMainGoFile(serviceName); err != nil {
		log.Fatalln(err)
	}

	//updateTransportHttp()
	if err := updateTransportHttpFile(serviceName); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s created successfully\n", serviceName)

	if err := goModTidy("./" + serviceName); err != nil {
		log.Fatalln(err)
	}
	if err := runGoImports("golang.org/x/tools/cmd/goimports@latest", "./"+serviceName); err != nil {
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
	if err := goModTidy("./" + serviceName + "_service"); err != nil {
		return err
	}
	if err := runGoImports("golang.org/x/tools/cmd/goimports@latest", "./"+serviceName+"_service"); err != nil {
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

func goModTidy(dir string) error {
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

func runGoImports(packagePath, dir string) error {
	if _, err := exec.LookPath("go"); err != nil {
		return errors.New("go path not found, please install go")
	} else {
		cmd := exec.Command("go", "install", packagePath)
		cmd.Dir = dir
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return errors.New(fmt.Sprintf("Failed to install goimports %v\n", err))
		}
	}

	if _, err := exec.LookPath("go"); err != nil {
		return errors.New("go path not found, please install go")
	} else {
		cmd := exec.Command("goimports", "-w", ".")
		cmd.Dir = dir
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return errors.New(fmt.Sprintf("Failed to execute goimports: %v\n", err))
		}
	}

	return nil
}

func updateMainGoFile(serviceName string) error {
	fSet := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fSet, "cmd/main.go", nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("error parsing file: %v", err)
	}

	if err := updateCreateHandler(parsedFile, serviceName); err != nil {
		return err
	}

	if err := updateImports(parsedFile, serviceName); err != nil {
		return err
	}

	outFile, err := os.Create("cmd/main.go")
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outFile.Close()

	if err := format.Node(outFile, fSet, parsedFile); err != nil {
		return fmt.Errorf("error formatting AST: %v", err)
	}

	return nil
}

func updateCreateHandler(file *ast.File, serviceName string) error {
	for _, decl := range file.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if ok && funcDecl.Name.Name == "createHandler" {
			for _, stmt := range funcDecl.Body.List {
				blockStmt, ok := stmt.(*ast.BlockStmt)
				if ok {
					newLine := fmt.Sprintf("%sHandler.NewHandler(group, %sService.NewService(db))", serviceName, serviceName)

					newExpr, err := parser.ParseExpr(strings.TrimSpace(newLine))
					if err != nil {
						return fmt.Errorf("error parsing new line: %v", err)
					}

					newStmt := &ast.ExprStmt{X: newExpr}

					blockStmt.List = append(blockStmt.List, newStmt)
				}
			}
		}
	}
	return nil
}

func updateImports(file *ast.File, newServiceName string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %v", err)
	}
	currentPackageName := filepath.Base(cwd)

	var importDecl *ast.GenDecl
	for _, decl := range file.Decls {
		if decl, ok := decl.(*ast.GenDecl); ok && decl.Tok == token.IMPORT {
			importDecl = decl
			break
		}
	}

	if importDecl != nil {
		import1 := &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("%sHandler \"%s/%s/handler\"", strings.ToLower(newServiceName), currentPackageName, newServiceName),
			},
		}
		import2 := &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("%sService \"%s/%s/service\"", strings.ToLower(newServiceName), currentPackageName, newServiceName),
			},
		}

		importDecl.Specs = append(importDecl.Specs, import1, import2)
	} else {
		return errors.New("import block not found in the file")
	}

	return nil
}

func updateTransportHttpFile(serviceName string) error {
	fSet := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fSet, "transport/service/http.go", nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("error parsing file: %v", err)
	}

	if err := updateNewService(parsedFile, serviceName); err != nil {
		return err
	}

	if err := updateImports(parsedFile, serviceName); err != nil {
		return err
	}

	outFile, err := os.Create("transport/service/http.go")
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outFile.Close()

	if err := format.Node(outFile, fSet, parsedFile); err != nil {
		return fmt.Errorf("error formatting AST: %v", err)
	}

	return nil
}

func updateNewService(file *ast.File, serviceName string) error {
	for _, decl := range file.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if ok && funcDecl.Name.Name == "NewService" {
			newLine := fmt.Sprintf("%sHandler.NewHandler(routerGroup, %sService.NewService(db))", serviceName, serviceName)

			newExpr, err := parser.ParseExpr(strings.TrimSpace(newLine))
			if err != nil {
				return fmt.Errorf("error parsing new line: %v", err)
			}

			newStmt := &ast.ExprStmt{X: newExpr}

			funcDecl.Body.List = append(funcDecl.Body.List, newStmt)
		}
	}
	return nil
}
