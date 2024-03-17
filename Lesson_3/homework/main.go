package main

import (
	"fmt"
	"homework3/Story"
)

func initCarriage(carriageNumberStart, carriageNumberHead, carriageNumberTail int) Carriage {
	carriageActions := make([]string, 7, 10)
	carriageActions[0] = "інформація для пасажирів"
	carriageActions[1] = "активувати стоп кран"
	carriageActions[2] = "зв'язок з провідником"
	carriageActions[3] = "перейти в наступний вагон"
	carriageActions[4] = "перейти в попередній вагон"

	var lastIndex int
	for i := len(carriageActions); i <= 0; i -= 1 {
		if carriageActions[i] != "" {
			lastIndex = i
			break
		}
	}
	carriageGame := Carriage{
		CarriageActions{
			Actions:    carriageActions,
			lastNonNil: lastIndex,
		},
		CarriageState{
			currentCarriageNumber: carriageNumberStart,
			carriageNumberHead:    carriageNumberHead,
			carriageNumberTail:    carriageNumberTail,
		},
	}
	return carriageGame
}

type CarriageState struct {
	currentCarriageNumber int
	carriageNumberHead    int
	carriageNumberTail    int
}

type CarriageActions struct {
	Actions    []string
	lastNonNil int
}

type Carriage struct {
	CarriageActions
	CarriageState
}

func (ca CarriageActions) showActions() {
	for actionInd, action := range ca.Actions {
		if action != "" {
			fmt.Println(actionInd, " ", action)
		}
	}
}

func (c *Carriage) goAnotherCarriage(direction int) {

	var directionSign string
	switch direction {
	case -1:
		directionSign = "<--"
	case 1:
		directionSign = "-->"
	default:
		fmt.Println("wrong direction value")
		return
	}

	fmt.Printf("Going %v carriage\n ", directionSign)
	carriageCandidate := c.currentCarriageNumber + direction

	switch {
	case carriageCandidate < c.carriageNumberHead:
		fmt.Println("Ви в хвості потягу")
	case carriageCandidate > c.carriageNumberTail:
		fmt.Println("Вперлись в вагон машиніста")
	default:
		Story.GoCarriageHeadAnimation(direction)
		c.currentCarriageNumber = carriageCandidate
	}
}

//func (cs *CarriageState) changeCurrentCarriage(newNumber int) {
//	cs.currentCarriageNumber = newNumber
//}

var (
	CarriageNumberStart int = 10
	carriageNumberHead  int = 1
	carriageNumberTail  int = 11
)

func main() {

	carriageGame := initCarriage(CarriageNumberStart, carriageNumberHead, carriageNumberTail)

	var input int
	for true {
		carriageGame.showActions()
		fmt.Scanf("%d", &input)
		carriageGame.goAnotherCarriage(input)
	}
}
