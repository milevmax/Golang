package routesAggregator

import (
	"homeworkTask1/primitives"
	"homeworkTask1/transport"
)

type PublicTransport interface {
	ReceivePassengers()
	DropOffPassengers()
}

func UseTransport(transport PublicTransport) {
	transport.ReceivePassengers()
	transport.DropOffPassengers()
}

func NewRouteTransport(transportName primitives.TransportType) PublicTransport {
	switch transportName {
	case primitives.Bus:
		t := transport.Bus{}
		return t
	case primitives.Train:
		t := transport.Train{}
		return t
	case primitives.Plane:
		t := transport.Plane{}
		return t
	default:
		return transport.ErrorTransport{}
	}
}

func GoOverRoute(route primitives.Route) {
	for _, t := range route {
		routeComponentTransport := NewRouteTransport(t)
		UseTransport(routeComponentTransport)
	}
}
