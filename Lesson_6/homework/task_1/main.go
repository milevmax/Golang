package main

import (
	"homeworkTask1/planing"
	"homeworkTask1/routesAggregator"
)

func main() {
	routeDetails := planing.Route{}
	routeDetails.AddTransportToRoute(planing.Bus)
	routeDetails.AddTransportToRoute(planing.Plane)
	routeDetails.ShowTransportsRoute()

	routesAggregator.GoOverRoute(routeDetails)

	routeDetails.AddTransportToRoute(planing.Train)
	routeDetails.AddTransportToRoute("Batmobile")
	routesAggregator.GoOverRoute(routeDetails)

}
