package matrix

type Pair struct {
	i, j int
}

func isMaxNumberInRow(n int, row []int) bool {
	for _, v := range row {
		if v > n {
			return false
		}
	}
	return true
}

func isMinNumberInColumn(n int, col []int) bool {
	for _, v := range col {
		if v < n {
			return false
		}
	}
	return true
}

func (m *Matrix) Saddle() []Pair {
	c := m.Cols()
	pair := []Pair{}

	for i, row := range m.Rows() {
		for j, value := range row {
			if isMaxNumberInRow(value, row) && isMinNumberInColumn(value, c[j]) {
				pair = append(pair, Pair{i, j})
			}
		}
	}
	return pair
}
