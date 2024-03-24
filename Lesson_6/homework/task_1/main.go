package main

import (
	"homeworkTask1/primitives"
	"homeworkTask1/routesAggregator"
)

func main() {
	routeDetails := primitives.Route{}
	routeDetails.AddTransportToRoute(primitives.Bus)
	routeDetails.AddTransportToRoute(primitives.Plane)
	routeDetails.ShowTransportsRoute()

	routesAggregator.GoOverRoute(routeDetails)

	routeDetails.AddTransportToRoute(primitives.Train)
	routeDetails.AddTransportToRoute(primitives.Train)

	routesAggregator.GoOverRoute(routeDetails)

}
