package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(test func(int) bool) Ints {
	var store Ints
	for _, n := range ints {
		if test(n) {
			store = append(store, n)
		}
	}
	return store
}

func (ints Ints) Discard(test func(int) bool) Ints {
	return ints.Keep(func(num int) bool {
		return !test(num)
	})
}

func (lists Lists) Keep(test func([]int) bool) Lists {
	var store Lists
	for _, list := range lists {
		if test(list) {
			store = append(store, list)
		}
	}
	return store
}

func (strings Strings) Keep(test func(string) bool) Strings {
	var store Strings
	for _, str := range strings {
		if test(str) {
			store = append(store, str)
		}
	}
	return store
}
