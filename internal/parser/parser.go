package parser

import (
	"os"
	"strings"

	"github.com/emicklei/proto"
)

const ignore = "dragoon:ignore"

type Service struct {
	Name string
	RPCs []RPC
}

type RPC struct {
	Name string
}

func Parse(basepath string) ([]Service, error) {
	reader, err := os.Open(basepath)
	if err != nil {
		return nil, err
	}

	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		return nil, err
	}

	var services []Service

	proto.Walk(definition, proto.WithService(func(s *proto.Service) {
		services = append(services, parseService(s))
	}))

	return services, nil
}

func parseService(s *proto.Service) Service {

	var rpcs []RPC

LOOP:
	for _, e := range s.Elements {

		switch t := e.(type) {
		case *proto.RPC:
			if t.Comment != nil {
				for _, l := range t.Comment.Lines {
					l := strings.TrimLeft(l, " ")
					if strings.HasPrefix(l, ignore) {
						continue LOOP
					}
				}
			}

			rpcs = append(rpcs, RPC{
				Name: t.Name,
			})
		}
	}

	return Service{
		Name: s.Name,
		RPCs: rpcs,
	}
}
