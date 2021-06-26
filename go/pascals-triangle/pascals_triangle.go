package pascal

func Triangle(n int) [][]int {
	arr := [][]int{}
	for i := 0; i < n; i++ {
		subArr := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || i == j { // checking if it is first or last index then set to 1
				subArr = append(subArr, 1)
			} else {
				add := arr[i-1][j-1] + arr[i-1][j] // Adding values with their previous index values & append to the current value
				subArr = append(subArr, add)
			}
		}
		arr = append(arr, subArr)
	}
	return arr
}
