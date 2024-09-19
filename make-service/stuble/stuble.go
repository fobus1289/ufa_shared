package stuble

import (
	"embed"
	"fmt"
	"io/fs"
	"path"
	"strings"
)

//go:embed **
var Templates embed.FS

//go:embed cmd.tpl
var Cmd string

//go:embed dto.tpl
var Dto string

//go:embed model.tpl
var Model string

//go:embed  service.tpl
var Service string

//go:embed handler.tpl
var Handler string

//go:embed http.tpl
var Http string

//go:embed gitignore.tpl
var Gitignore string

//go:embed env.tpl
var Env string

//go:embed README.tpl
var README string

func ReadDirs(dirname string, rec ...bool) ([]string, error) {

	var names []string
	{
		dirs, err := Templates.ReadDir(dirname)
		{
			if err != nil {
				return nil, err
			}
		}

		for _, dir := range dirs {
			if dir.IsDir() {
				names = append(names, dir.Name())
				if len(rec) > 0 {
					ff, _ := ReadDirs(dir.Name(), rec...)
					names = append(names, ff...)
				}
			}
		}
	}

	return names, nil
}

func WalkEmbedFS(fsys fs.FS, dir, serviceName string) error {
	entries, err := fs.ReadDir(fsys, dir)
	{
		if err != nil {
			return err
		}
	}

	for _, entry := range entries {
		fullPath := path.Join(dir, entry.Name())

		transformFullname := strings.Replace(fullPath, "[service]", serviceName, -1)

		if entry.IsDir() {
			fmt.Printf("Директория: %s\n", transformFullname)
			if err := WalkEmbedFS(fsys, fullPath, serviceName); err != nil {
				return err
			}
		} else {
			fmt.Printf("Файл: %s\n", transformFullname)

			// Если вам нужно прочитать содержимое файла:
			content, err := fs.ReadFile(fsys, fullPath)
			if err != nil {
				return err
			}
			fmt.Printf("  Содержимое (%d байт): %s\n", len(content), string(content[:min(len(content), 50)]))
		}
	}

	return nil
}

func WalkPrintFS(fsys fs.FS, dir, serviceName, pwd string, depth int) error {
	entries, err := fs.ReadDir(fsys, dir)
	{
		if err != nil {
			return err
		}
	}

	for i, entry := range entries {
		fullPath := path.Join(dir, entry.Name())
		transformFullname := strings.Replace(fullPath, "[service]", serviceName, -1)
		transformFullname = path.Join(strings.Replace(transformFullname, "soa", pwd, -1))
		prefix := strings.Repeat("  ", depth)

		isLast := i == len(entries)-1
		var marker string
		if isLast {
			marker = "└──"
		} else {
			marker = "├──"
		}

		if entry.IsDir() {
			fmt.Printf("%s%s 📁 %s\n", prefix, marker, transformFullname)
			err := WalkPrintFS(fsys, fullPath, serviceName, pwd, depth+1)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("%s%s 📄 %s\n", prefix, marker, transformFullname)
		}
	}

	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
