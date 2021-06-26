package grains

import "errors"

func Total() uint64 {
	return 1<<64 - 1
}

func Square(squareNumber int) (uint64, error) {
	if squareNumber >= 1 && squareNumber <= 64 {
		return (1 << uint(squareNumber-1)), nil
	}
	return 0, errors.New("Square does not exists")
}
