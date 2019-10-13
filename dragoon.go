package dragoon

type Service struct {
	Name string
	RPCs []RPC
}

type RPC struct {
	Name string
}
