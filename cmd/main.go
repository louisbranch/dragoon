package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/luizbranco/dragoon"
	"github.com/luizbranco/dragoon/internal/dot"
	"github.com/luizbranco/dragoon/internal/parser"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("missing file(s) path\n")
		os.Exit(1)
	}

	var files []string

	for i := 1; i < len(os.Args); i++ {

		root := os.Args[i]

		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("failed to open file %q - %s", filepath.Join(root, path), err)
				os.Exit(1)
			}

			if filepath.Ext(path) == ".proto" {
				files = append(files, path)
			}

			return nil
		})

	}

	var services []dragoon.Service

	for _, f := range files {

		s, err := parser.Parse(f)
		if err != nil {
			fmt.Printf("failed to parse file %q - %s", f, err)
			os.Exit(1)
		}

		services = append(services, s...)
	}

	graph := dot.Graph(services)

	fmt.Println(graph)
}
