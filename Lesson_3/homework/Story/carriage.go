package Story

import "fmt"

func InitCarriage(carriageNumberStart, carriageNumberHead, carriageNumberTail int) Carriage {
	carriageActions := make([]string, 7, 10)
	carriageActions[0] = "інформація для пасажирів"
	carriageActions[1] = "активувати стоп кран"
	carriageActions[2] = "зв'язок з провідником"
	carriageActions[3] = "перейти в інший вагон"

	var lastIndex int
	//for i := len(carriageActions); i <= 0; i -= 1 {
	//	if carriageActions[i] != "" {
	//		lastIndex = i
	//		break
	//	}
	//}
	lastIndex = 3

	carriageGame := Carriage{
		CarriageActions{
			Actions:       carriageActions,
			lastNonNilInd: lastIndex,
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
	Actions       []string
	lastNonNilInd int
	lastAction    int
}

type Carriage struct {
	CarriageActions
	CarriageState
}

func (ca CarriageActions) ShowActions() {
	for actionInd, action := range ca.Actions {
		if action != "" {
			fmt.Println(actionInd, " ", action)
		}
	}
}

func (c *Carriage) TakeInput() int {
	c.ShowActions()
	var input int
	fmt.Print("Обери номер дії: ")
	fmt.Scanf("%d", &input)
	c.lastAction = input
	return input
}

func (c *Carriage) HandleInput(inputCarriage int) (updatedState string) {
	switch inputCarriage {
	case 0:
		updatedState = InfoBoardStateName
	case 1:
		updatedState = StopTapName
	case 2:
		updatedState = CommunicationDeviceName
	case 3:
		updatedState = CarriageMovingStateName
	case 4:
		updatedState = EavesdropStateName
	}
	return updatedState
}

func (c *Carriage) TakeInputMovement() int {
	var input int
	fmt.Print("Пройти в наступний вагон - 1\nПройти в попередній вагон - 2\nНазад - 3\n--:")
	fmt.Scanf("%d", &input)
	if input == 2 {
		input = -1
	}
	return input
}

func (c *Carriage) GoAnotherCarriage(direction int) {

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

	fmt.Printf("Проходжу %v вагон\n ", directionSign)
	carriageCandidate := c.currentCarriageNumber + direction

	switch {
	case carriageCandidate < c.carriageNumberHead:
		fmt.Println("Ви в хвості потягу")
	case carriageCandidate > c.carriageNumberTail:
		fmt.Println("Вперлись в вагон машиніста")
	default:
		GoCarriageHeadAnimation(direction)
		c.currentCarriageNumber = carriageCandidate
	}
	c.UpdateActions()
}

func (c *Carriage) UpdateActions() {
	if c.currentCarriageNumber == c.carriageNumberTail && c.Actions[c.lastNonNilInd] != "підслухати кабіну провідника" {
		c.lastNonNilInd += 1
		c.Actions[c.lastNonNilInd] = "підслухати кабіну провідника"
	} else {
		if c.Actions[c.lastNonNilInd] == "підслухати кабіну провідника" {
			c.Actions[c.lastNonNilInd] = ""
			c.lastNonNilInd -= 1
		}
	}
}
