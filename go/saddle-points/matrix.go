package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	values     []int
	columnSize int
}

func (m Matrix) Cols() [][]int {
	newColumn := make([][]int, m.columnSize)
	rSize := len(m.values) / m.columnSize

	for i := 0; i < m.columnSize; i++ {
		newColumn[i] = make([]int, rSize)
		for j := 0; j < rSize; j++ {
			newColumn[i][j] = m.values[j*m.columnSize+i]
		}
	}

	return newColumn
}

func (m Matrix) Rows() [][]int {
	rSize := len(m.values) / m.columnSize
	newRow := make([][]int, rSize)

	for i := 0; i < rSize; i++ {
		newRow[i] = make([]int, m.columnSize)
		for j := 0; j < m.columnSize; j++ {
			newRow[i][j] = m.values[i*m.columnSize+j]
		}
	}

	return newRow
}

func (m *Matrix) Set(row, col, value int) bool {
	rSize := len(m.values) / m.columnSize
	if row < 0 || row >= rSize || col < 0 || col >= m.columnSize {
		return false
	}
	if row*m.columnSize+col >= len(m.values)+1 {
		return false
	}

	m.values[row*m.columnSize+col] = value
	return true
}

func New(str string) (*Matrix, error) {
	m := &Matrix{}
	arr := strings.Split(str, "\n")
	for _, v := range arr {
		strArr := strings.Fields(v)
		if len(strArr) == 0 {
			return nil, errors.New("Empty Array !")
		}

		if m.columnSize == 0 {
			m.columnSize = len(strArr)
		} else if m.columnSize != len(strArr) {
			return nil, errors.New("Column Not matched !")
		}

		for _, value := range strArr {
			i, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return nil, err
			}
			m.values = append(m.values, int(i))
		}
	}

	return m, nil
}
