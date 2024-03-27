package routesAggregator

import (
	"fmt"
	"homeworkTask1/planing"
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

func NewRouteTransport(transportName planing.TransportType) PublicTransport {
	switch transportName {
	case planing.Bus:
		t := transport.Bus{}
		return t
	case planing.Train:
		t := transport.Train{}
		return t
	case planing.Plane:
		t := transport.Plane{}
		return t
	default:
		return transport.ErrorTransport{ErrorTransportName: transportName}
	}
}

func GoOverRoute(route planing.Route) {
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
