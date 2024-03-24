package transport

import "fmt"

type Bus struct {
}

func (b Bus) ReceivePassengers() {
	fmt.Println("Bus fulled by passengers")
}
func (b Bus) DropOffPassengers() {
	fmt.Println("Bus dropped passengers")
}

type Plane struct {
}

func (p Plane) ReceivePassengers() {
	fmt.Println("Plane fulled by passengers")
}
func (p Plane) DropOffPassengers() {
	fmt.Println("Plane dropped passengers")
}

type Train struct {
}

func (t Train) ReceivePassengers() {
	fmt.Println("Train fulled by passengers")
}
func (t Train) DropOffPassengers() {
	fmt.Println("Train dropped passengers")
}

type ErrorTransport struct {
}

func (t ErrorTransport) ReceivePassengers() {
	fmt.Println("ErrorTransport")
}
func (t ErrorTransport) DropOffPassengers() {
	fmt.Println("ErrorTransport")
}
