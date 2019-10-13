package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/emicklei/proto"
)

func main() {
	reader, _ := os.Open(os.Args[1])
	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	proto.Walk(definition, proto.WithService(handleService))
}

func handleService(s *proto.Service) {
	fmt.Println(s.Name)

LOOP:
	for _, e := range s.Elements {
		switch t := e.(type) {
		case *proto.RPC:
			if t.Comment != nil {
				for _, l := range t.Comment.Lines {
					l := strings.TrimLeft(l, " ")
					if strings.HasPrefix(l, "dragon:ignore") {
						continue LOOP
					}
				}
			}
			fmt.Printf("\t%s\n", t.Name)
		case *proto.Comment:
		default:
			fmt.Printf("%+v", e)
		}
	}
}
