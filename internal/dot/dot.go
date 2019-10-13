package dot

import (
	"github.com/emicklei/dot"
	"github.com/luizbranco/dragoon"
)

func Graph(services []dragoon.Service) *dot.Graph {
	graph := dot.NewGraph(dot.Directed)

	for _, s := range services {
		cluster := graph.Subgraph(s.Name, dot.ClusterOption{})

		for _, rpc := range s.RPCs {
			cluster.Node(rpc.Name)
		}
	}

	return graph
}
