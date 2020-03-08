package calculator

const price = 100.0

func CalBooksPrice(books []int) float64 {
	if len(books) == 2 {
		return price * 2 * 0.95
	}
	if len(books) == 3 {
		return price * 3 * 0.9
	}
	if len(books) == 4 {
		return price * 4 * 0.8
	}
	if len(books) == 5 {
		return price * 5 * 0.75
	}

	return price
}
