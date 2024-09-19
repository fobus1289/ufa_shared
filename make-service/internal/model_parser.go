package internal

import (
	"fmt"

	"golang.org/x/mod/modfile"
)

func ParseMod(content string) (*modfile.File, error) {

	modfile, err := modfile.Parse("_", []byte(content), nil)
	{
		if err != nil {
			return nil, fmt.Errorf("error parsing go.mod: %v", err)
		}
	}

	return modfile, nil
}
