package series

func UnsafeFirst(n int, s string) string {
	first, _ := First(n, s)
	return first
}

func All(n int, s string) []string {
	result := []string{}
	for i := 0; i+n <= len(s); i++ {
		result = append(result, string(s[i:i+n]))
	}
	return result
}

func First(n int, s string) (first string, ok bool) {
	if len(s) >= n {
		first, ok = s[:n], true
	}
	return
}
