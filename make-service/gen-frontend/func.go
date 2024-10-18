package genfrontend

import (
	"encoding/json"
	"strings"
)

func getTitle() string {
	return spec.GetInfo().Title
}

func getPaths(tags ...string) map[string]any {

	var (
		paths  = make(map[string]PathItem)
		anyMap = make(map[string]any)
	)

	if selectedModules == nil {
		paths = spec.Paths
	} else {
		for _, selectedPath := range selectedModules {
			for path, v := range spec.Paths {
				if selectedPath == strings.Split(path, "/")[1] {
					paths[path] = v
				}
			}
		}
	}

	byt, err := json.Marshal(paths)
	{
		if err != nil {
			panic(err)
		}
	}

	err = json.Unmarshal(byt, &anyMap)
	{
		if err != nil {
			panic(err)
		}
	}

	// Output to verify
	return anyMap
}
