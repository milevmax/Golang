package transport

import (
	"fmt"
)

type TransportTechDetails struct {
	AvgSpeed    int
	MaxDistance int
}

func (td TransportTechDetails) GetMaxDistance() int {
	return td.MaxDistance
}

func (td TransportTechDetails) ToBeOnTheRoad(distance int) {
	timeInRoad := float32(distance) / float32(td.AvgSpeed)
	printHoursAndMinutes(timeInRoad)
}

type Bus struct {
	RoadNumber  int
	TicketPrice int
	TransportTechDetails
}

func (b Bus) ReceivePassengers() {
	fmt.Println("Bus fulled by passengers")
}

func (b Bus) DropOffPassengers() {
	fmt.Println("Bus dropped passengers")
}

type Plane struct {
	flightNumber int
	meal         bool
	TransportTechDetails
}

func (p Plane) ReceivePassengers() {
	fmt.Println("Plane fulled by passengers")
}

func (p Plane) DropOffPassengers() {
	fmt.Println("Plane dropped passengers")
}

type Train struct {
	trainNumber    int
	carriageNumber int
	TransportTechDetails
}

func (t Train) ReceivePassengers() {
	fmt.Println("Train fulled by passengers")
}

func (t Train) DropOffPassengers() {
	fmt.Println("Train dropped passengers")
}
