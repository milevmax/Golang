package primitives

import "fmt"

type TransportType string

const (
	Bus   TransportType = "bus"
	Train TransportType = "train"
	Plane TransportType = "plane"
)

type Route []TransportType

func (r *Route) AddTransportToRoute(transport TransportType) {
	*r = append(*r, transport)
}

func (r *Route) ShowTransportsRoute() {
	for i, v := range *r {
		fmt.Printf("%d transport: %s\n", i+1, v)
	}
}
