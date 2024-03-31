package routes

import (
	"errors"
	"fmt"
	"homeworkTask1/transport"
)

type PublicTransport interface {
	ReceivePassengers()
	ToBeOnTheRoad(distance int)
	DropOffPassengers()
}

type transportDetails interface {
	GetMaxDistance() int
}

type PublicTransportDetails interface {
	PublicTransport
	transportDetails
}

func NewPublicTransport(transportName TransportType) (PublicTransportDetails, error) {
	switch transportName {
	case Bus:
		t := transport.Bus{TransportTechDetails: transport.TransportTechDetails{AvgSpeed: BusAvgSpeed,
			MaxDistance: BusMaxDistance}}
		return t, nil
	case Train:
		t := transport.Train{TransportTechDetails: transport.TransportTechDetails{AvgSpeed: TrainAvgSpeed,
			MaxDistance: TrainMaxDistance}}
		return t, nil
	case Plane:
		t := transport.Plane{TransportTechDetails: transport.TransportTechDetails{AvgSpeed: PlaneAvgSpeed,
			MaxDistance: PlaneMaxDistance}}
		return t, nil
	default:
		return nil, errors.New("invalid transport type")
	}
}

type TripComponent struct {
	TravelDistance int
	MovingMethod   PublicTransport
}

func validateRoute(td transportDetails, distance int) bool {
	return td.GetMaxDistance() > distance
}

type Trip []TripComponent

func (t *Trip) AddTripComponent(transport PublicTransportDetails, distance int) {
	if validateRoute(transport, distance) {
		tr := TripComponent{distance, transport}
		*t = append(*t, tr)
	} else {
		fmt.Printf("Distance is too far for chosen transport. Max possible range: %d\n", transport.GetMaxDistance())
	}
}

func (t *Trip) PassTrip() {
	for _, tripComponent := range *t {
		tripComponent.MovingMethod.ReceivePassengers()
		tripComponent.MovingMethod.ToBeOnTheRoad(tripComponent.TravelDistance)
		tripComponent.MovingMethod.DropOffPassengers()
	}
}
