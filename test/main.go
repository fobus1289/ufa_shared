package main

import (
	"log"

	"github.com/fobus1289/ufa_shared/utils"
)

func main() {
	log.Println(utils.CurrentDir())
	log.Println(utils.CurrentFile())
}
