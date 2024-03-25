package routesAggregator

import (
	"fmt"
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
	fmt.Println("\nTransfer to another transport\n")
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
		return transport.ErrorTransport{ErrorTransportName: transportName}
	}
}

func GoOverRoute(route primitives.Route) {
	for _, t := range route {
		routeComponentTransport := NewRouteTransport(t)
		if _, ok := routeComponentTransport.(transport.ErrorTransport); ok {
			fmt.Println("unexpected transport!")
			fmt.Println(routeComponentTransport)
			return
		}
		UseTransport(routeComponentTransport)
	}
}
