package main

import (
	"fmt"
	"homeworkTask1/routes"
)

func main() {

	busTransport1, _ := routes.NewPublicTransport(routes.Bus)
	planeTransport1, _ := routes.NewPublicTransport(routes.Plane)
	busTransport2, _ := routes.NewPublicTransport(routes.Bus)
	trainTransport1, _ := routes.NewPublicTransport(routes.Train)

	fmt.Println(busTransport1.GetMaxDistance(), planeTransport1.GetMaxDistance())

	trip := routes.Trip{}
	trip.AddTripComponent(busTransport1, 501)
	trip.AddTripComponent(planeTransport1, 2000)
	trip.AddTripComponent(busTransport2, 120)
	trip.AddTripComponent(trainTransport1, 530)

	trip.PassTrip()
}
