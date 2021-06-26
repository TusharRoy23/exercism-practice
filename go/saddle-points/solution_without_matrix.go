import (
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	i, j int
}

func findLargestNumberFromRow(arr []int) int {
	max := arr[0]
	indexNumber := 0
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			indexNumber = i
		}
	}
	return indexNumber
}

func compareNumberWithinColumn(selectedNumber int, index int, mainArr [][]int) bool {
	status := true
	for i := 0; i < len(mainArr); i++ {
		if mainArr[i][index] < selectedNumber {
			status = false
			break
		}
	}
	return status
}

func main() {
	str := "4 5 4\n3 5 5\n1 5 4"
	mainArr := [][]int{}
	pair := []Pair{}
	arr := strings.Split(str, "\n")
	for _, v := range arr {
		value := []int{}

		for _, v := range strings.Fields(v) {
			i, _ := strconv.ParseInt(v, 10, 32)
			value = append(value, int(i))
		}
		mainArr = append(mainArr, value)
	}

	for i := 0; i < len(mainArr); i++ {
		index := findLargestNumberFromRow(mainArr[i])
		isSmallestNumber := compareNumberWithinColumn(mainArr[i][index], index, mainArr)
		if isSmallestNumber {
			pair = append(pair, Pair{i, index})
		}
	}

	fmt.Println(pair)
}
