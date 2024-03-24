package main

import (
	"homework3/Story"
)

//func (cs *CarriageState) changeCurrentCarriage(newNumber int) {
//	cs.currentCarriageNumber = newNumber
//}

var (
	CarriageNumberStart int    = 10
	carriageNumberHead  int    = 1
	carriageNumberTail  int    = 11
	playerState         string = Story.IntroStateName
)

// func init game func to create intro, carriage, infoboard, stop tap, communication device

func main() {
	introState := Story.InitIntro()
	carriageState := Story.InitCarriage(CarriageNumberStart, carriageNumberHead, carriageNumberTail)
	eavedropState := Story.InitEavedrop()

	for playerState != "gg" {
		switch playerState {
		case Story.IntroStateName:
			introState.TakeInput()
			playerState = introState.HandleInput()
		case Story.CarriageStateName:
			inputCarriage := carriageState.TakeInput()
			playerState = carriageState.HandleInput(inputCarriage)
		case Story.CarriageMovingStateName:
			inputCarriageMoving := carriageState.TakeInputMovement()
			if inputCarriageMoving == 1 || inputCarriageMoving == -1 {
				carriageState.GoAnotherCarriage(inputCarriageMoving)
			} else {
				playerState = Story.CarriageStateName
			}
		case Story.EavesdropStateName:
			eavedropState.Listen()
			playerState = Story.CarriageStateName
		default:
			playerState = "gg"
		}
	}
}

//
//func main() {
//
//	carriageGame := Story.InitCarriage(CarriageNumberStart, carriageNumberHead, carriageNumberTail)
//
//	var input int
//	for true {
//		carriageGame.ShowActions()
//		fmt.Scanf("%d", &input)
//		carriageGame.GoAnotherCarriage(input)
//	}
//}
