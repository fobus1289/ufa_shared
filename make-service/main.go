package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fobus1289/ufa_shared/make-service/internal"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		projectName := promptInput("Enter project name: ")
		modPath := promptInput("Enter project mod path: ")
		switch os.Args[1] {
		case "--new":
			internal.NewService(projectName, modPath)
		case "--add":
			internal.AddService(projectName, modPath)
		default:
			log.Fatalln(errors.New("unknown flag"))
		}
	} else {
		log.Fatalln(errors.New("not enough arguments"))
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
