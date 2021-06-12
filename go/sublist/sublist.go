package sublist

type Relation string

func returnMatchedResults(a int, b int) Relation {
	if a == b {
		return Relation("equal")
	} else if a > b {
		return Relation("superlist")
	} else {
		return Relation("sublist")
	}
}

func compareSlice(bigger []int, smaller []int, lengthOfListOne int, lengthOfListTwo int) Relation {
	startingIndex := 0

	for i := 0; i < len(bigger); i++ {
		con := 0
		if len(smaller)+i > len(bigger) {
			return Relation("unequal")
		}
		newSlice := bigger[startingIndex : len(smaller)+i]
		for j := 0; j < len(newSlice); j++ {
			if smaller[j] == newSlice[j] {
				con++
			}
			if con == len(smaller) {
				return returnMatchedResults(lengthOfListOne, lengthOfListTwo)
			}
		}
		startingIndex++
	}
	return Relation("unequal")
}

func Sublist(listOne []int, listTwo []int) Relation {
	if len(listOne) == 0 && len(listTwo) == 0 {
		return Relation("equal")
	} else if len(listOne) == 0 && len(listTwo) > 0 {
		return Relation("sublist")
	} else if len(listTwo) == 0 && len(listOne) > 0 {
		return Relation("superlist")
	} else if len(listOne) >= len(listTwo) {
		return compareSlice(listOne, listTwo, len(listOne), len(listTwo))
	} else {
		return compareSlice(listTwo, listOne, len(listOne), len(listTwo))
	}
}
