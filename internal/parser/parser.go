package parser

import (
	"os"
	"strings"

	"github.com/emicklei/proto"
	"github.com/luizbranco/dragoon"
)

const ignore = "dragoon:ignore"

func Parse(basepath string) ([]dragoon.Service, error) {
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

	var services []dragoon.Service

	proto.Walk(definition, proto.WithService(func(s *proto.Service) {
		services = append(services, parseService(s))
	}))

	return services, nil
}

func parseService(s *proto.Service) dragoon.Service {

	var rpcs []dragoon.RPC

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

			rpcs = append(rpcs, dragoon.RPC{
				Name: t.Name,
			})
		}
	}

	return dragoon.Service{
		Name: s.Name,
		RPCs: rpcs,
	}
}
