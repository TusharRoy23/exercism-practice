package allergies

var allergens = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func AllergicTo(score uint, substance string) bool {
	for _, allergen := range Allergies(score) {
		if allergen == substance {
			return true
		}
	}
	return false
}

func Allergies(code uint) []string {
	allergies := []string{}

	for _, a := range allergens {
		if code <= 0 {
			break
		}

		if code%2 == 1 {
			allergies = append(allergies, a)
		}

		code = code >> 1
	}

	return allergies
}
