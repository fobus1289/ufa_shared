package service

import (
	"bytes"
	"errors"
	"fmt"

	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func CreateFiles(serviceName, modPath string, files map[string]string) error {
	for filePath, content := range files {

		file, err := os.Create(filePath)
		if err != nil {
			return errors.New(fmt.Sprintf("create file error %v", err))
		}

		m := map[string]string{
			"ServiceName": serviceName,
			"ModPath":     modPath,
		}

		var buffer bytes.Buffer
		{
			if err := Tmp(content).Execute(&buffer, m); err != nil {
				_ = file.Close()
				return errors.New(fmt.Sprintf("content copy error %v", err))
			}

			if _, err := file.Write(buffer.Bytes()); err != nil {
				_ = file.Close()
				return errors.New(fmt.Sprintf("write content error %v", err))
			}

			_ = file.Close()
		}
	}

	return nil
}

func UpdateMainGoFile(serviceName, modPath string) error {
	fSet := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fSet, "cmd/main.go", nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("error parsing file: %v", err)
	}

	if err := UpdateCreateHandler(parsedFile, serviceName); err != nil {
		return err
	}

	if err := UpdateImports(parsedFile, serviceName, modPath); err != nil {
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

func UpdateCreateHandler(file *ast.File, serviceName string) error {
	for _, decl := range file.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if ok && funcDecl.Name.Name == "createHandler" {
			for _, stmt := range funcDecl.Body.List {
				blockStmt, ok := stmt.(*ast.BlockStmt)
				if ok {
					newLine := fmt.Sprintf("%sHandler.NewHandler(group, %sService.NewService(db))", ToLowerCamel(serviceName), ToLowerCamel(serviceName))

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

func UpdateTransportHttpFile(serviceName, modPath string) error {
	fSet := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fSet, "transport/service/http.go", nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("error parsing file: %v", err)
	}

	if err := UpdateNewService(parsedFile, serviceName); err != nil {
		return err
	}

	if err := UpdateImports(parsedFile, serviceName, modPath); err != nil {
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

func UpdateNewService(file *ast.File, serviceName string) error {
	for _, decl := range file.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if ok && funcDecl.Name.Name == "NewService" {
			newLine := fmt.Sprintf("%sHandler.NewHandler(routerGroup, %sService.NewService(db))", ToLowerCamel(serviceName), ToLowerCamel(serviceName))

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
