package dot

import (
	"testing"

	"github.com/luizbranco/dragoon"
)

func TestGraph(t *testing.T) {
	tcs := map[string]struct {
		services []dragoon.Service
		output   string
	}{
		"ok": {
			services: []dragoon.Service{
				{
					Name: "Balancer",
					RPCs: []dragoon.RPC{
						{Name: "CreditAccount"},
						{Name: "DebitAccount"},
					},
				},
			},
			output: `digraph  {
	subgraph cluster_s0 {
		ID = "cluster_s0";
		label="Balancer";
		n1[label="CreditAccount"];
		n2[label="DebitAccount"];
		
	}
	
}`,
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {

			graph := Graph(tc.services)

			if graph.String() != tc.output {
				t.Errorf("expected\n%q\n%q", tc.output, graph.String())
			}

		})
	}
}
