package dot

import (
	"github.com/emicklei/dot"
	"github.com/luizbranco/dragoon"
)

func Graph(services []dragoon.Service) *dot.Graph {
	graph := dot.NewGraph(dot.Directed)

	var firsts []dot.Node

	for _, s := range services {
		cluster := graph.Subgraph(s.Name, dot.ClusterOption{})

		for i, rpc := range s.RPCs {
			n := cluster.Node(rpc.Name)

			if i == 0 {
				firsts = append(firsts, n)
			}
		}
	}

	for i, n := range firsts {
		if i == len(firsts)-1 {
			break
		}

		e := n.Edge(firsts[i+1])
		e.Attr("style", "invis")
	}

	return graph
}
