package secret

func doReverse(code []string) []string {
	revArr := []string{}
	for i := len(code) - 1; i >= 0; i-- {
		revArr = append(revArr, code[i])
	}
	return revArr
}

func Handshake(n uint) []string {
	codes := []string{}
	if n&1 == 1 {
		codes = append(codes, "wink")
	}
	if n&2 == 2 {
		codes = append(codes, "double blink")
	}

	if n&4 == 4 {
		codes = append(codes, "close your eyes")
	}

	if n&8 == 8 {
		codes = append(codes, "jump")
	}

	if n&16 == 16 {
		codes = doReverse(codes)
	}
	return codes
}
