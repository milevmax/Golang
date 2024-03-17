package Story

import (
	"fmt"
	"time"
)

//TODO: union intro and eavedrop struct

func InitEavedrop() Eavedrop {
	EavedropGame := Eavedrop{EavedropText: "\n- Чуєх, Льоха, а Міхалич стоп кран полагодив?\n- Той який чомусь напругу з дверей прибирає?\n- Ага, в третім вагоні\n- Ні, сказав з відпустки прийде і зробе\n- Як завжди...\n"}
	return EavedropGame
}

type Eavedrop struct {
	EavedropText     string
	EavedropListened bool
}

func (e *Eavedrop) Listen() {
	fmt.Println(e.EavedropText)
	time.Sleep(2 * time.Second)
	e.EavedropText = "..."
}
