package main

import (
	"errors"
	"github.com/fobus1289/ufa_shared/make-service/internal"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		switch os.Args[1] {
		case "--new":
			internal.NewService(os.Args[2])
		case "--add":
			internal.AddService(os.Args[2])
		default:
			log.Fatalln(errors.New("unknown flag"))
		}
	} else {
		log.Fatalln(errors.New("not enough arguments"))
	}
}
