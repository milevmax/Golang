package routes

type TransportType string

const (
	Bus            TransportType = "bus"
	BusMaxDistance int           = 500
	BusAvgSpeed                  = 80

	Train            TransportType = "train"
	TrainMaxDistance int           = 1000
	TrainAvgSpeed                  = 50

	Plane            TransportType = "plane"
	PlaneMaxDistance int           = 5000
	PlaneAvgSpeed    int           = 950
)
