package flatten

func checkingArr(arr interface{}, result *[]interface{}) {
	for _, value := range arr.([]interface{}) {
		if value != nil {
			switch con := value.(type) {
			case int:
				*result = append(*result, value)
			case []interface{}:
				checkingArr(con, result)
			}
		}
	}
}

func Flatten(arr interface{}) []interface{} {
	result := []interface{}{}
	checkingArr(arr, &result)
	return result
}
