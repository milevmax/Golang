package Story

import (
	"fmt"
	"strings"
	"time"
)

var (
	IntroStateName          string = "intro"
	CarriageStateName       string = "carriage"
	InfoBoardStateName      string = "info board"
	CommunicationDeviceName string = "communication device"
	StopTapName             string = "stop tap"
	CarriageMovingStateName string = "carriage move"
	EavesdropStateName      string = "eavesdrop"
)

func ReplaceByIndex(sourceStr string, ind int, replacement rune) (string, rune) {
	symbols := []rune(sourceStr)
	oldValue := symbols[ind]
	symbols[ind] = replacement
	return string(symbols), oldValue
}

func GoCarriageHeadAnimation(direction int) {

	var (
		carriageStringVisualisation string = "|-----------||-----*-----||-----------|"
		playerPosition              int    = 19
		playerPointVisualisation    rune   = '*'
		directionComponent          int    = direction
		oldVal                      rune   = '-'
		destinationPosition         int    = playerPosition + directionComponent*15
		indReplaceNext              int
		indReplacePrev              int
	)

	time.Sleep(time.Second)
	for ind := playerPosition + directionComponent; ind != destinationPosition; ind += directionComponent {

		if indReplaceNext == 0 {
			indReplacePrev = ind - directionComponent
		} else {
			indReplacePrev = indReplaceNext
		}

		if carriageStringVisualisation[ind] == uint8('|') {
			continue
		}
		indReplaceNext = ind
		fmt.Print(carriageStringVisualisation)

		if ind == destinationPosition-directionComponent {
			time.Sleep(2 * time.Second)
		} else {
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Printf("\r%s\r", strings.Repeat(" ", len(carriageStringVisualisation)+1))

		carriageStringVisualisation, _ = ReplaceByIndex(carriageStringVisualisation, indReplacePrev, oldVal)
		carriageStringVisualisation, _ = ReplaceByIndex(carriageStringVisualisation, indReplaceNext, playerPointVisualisation)
	}
	fmt.Println()
}
