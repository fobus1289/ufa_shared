package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	gentest "github.com/fobus1289/ufa_shared/make-service/gen-test"
	"github.com/fobus1289/ufa_shared/make-service/internal"
)

func main() {

	if len(os.Args) == 0 {
		log.Fatalln(errors.New("not enough arguments"))
	}

	switch os.Args[1] {
	case "--new":
		projectName := promptInput("Enter project name: ")
		modPath := promptInput("Enter project mod path: ")
		internal.NewService(projectName, modPath)
	case "--add":
		projectName := promptInput("Enter project name: ")
		modPath := promptInput("Enter project mod path: ")
		internal.AddService(projectName, modPath)
	case "--test":
		swaggFilePath := promptInput("Enter swagg file path: ")
		testPath := promptInput("Enter test path: ")
		gentest.GenerateTest(swaggFilePath, testPath)
	default:
		log.Fatalln(errors.New("unknown flag"))
	}
}

func promptInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(input)
}
