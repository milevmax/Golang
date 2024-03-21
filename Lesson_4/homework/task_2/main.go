package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
)

type MyStruct struct {
	Id int
}

func SliceMax(structures []MyStruct) MyStruct {
	MaxValStruct := structures[0]
	for _, s := range structures {
		if s.Id > MaxValStruct.Id {
			MaxValStruct = s
		}
	}
	return MaxValStruct
}

func SliceToSet(structures []MyStruct) (setStructures []MyStruct) {

	setStructures = make([]MyStruct, 0, len(structures))

	var (
		LenInput     int      = len(structures)
		maxSorted    MyStruct = MyStruct{Id: -1}
		minUnsortedS MyStruct = SliceMax(structures)
	)

	for i := 0; i < LenInput; i++ {
		minUnsortedIter := minUnsortedS
		for _, s := range structures {
			if s.Id > maxSorted.Id && s.Id < minUnsortedIter.Id {
				minUnsortedIter = s
			}
		}
		if minUnsortedIter.Id == maxSorted.Id {
			return setStructures
		}
		setStructures = append(setStructures, minUnsortedIter)
		maxSorted = minUnsortedIter
	}
	return setStructures
}

func SliceToSet_v2(structures []MyStruct) (setStructures []MyStruct) {
	setStructures = make([]MyStruct, 0, len(structures))
	sort.SliceStable(structures, func(i, j int) bool {
		return structures[i].Id < structures[j].Id
	})

	var lastValue MyStruct = MyStruct{Id: -1}
	for _, v := range structures {
		if v == lastValue {
			continue
		}
		setStructures = append(setStructures, v)
		lastValue = v
	}
	return setStructures
}

func main() {
	lenStruct := 20
	workStruct := [20]MyStruct{}

	for i, _ := range workStruct {
		structIter := MyStruct{Id: rand.IntN(lenStruct / 3)}
		workStruct[i] = structIter
	}
	fmt.Println(workStruct)
	inputFunc := workStruct[0:6]
	fmt.Println(inputFunc)

	funcRes := SliceToSet(inputFunc)
	fmt.Println(funcRes)

	funcRes2 := SliceToSet_v2(inputFunc)
	fmt.Println(funcRes2)

}
