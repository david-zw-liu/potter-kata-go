package calculator

const price = 100.0

func getPriceByDiffCount(diffCount int) float64 {
	switch diffCount {
	case 2:
		return price * 2 * 0.95
	case 3:
		return price * 3 * 0.9
	case 4:
		return price * 4 * 0.8
	case 5:
		return price * 5 * 0.75
	default:
		return price
	}
}

func buildBookCount(books []int) map[int]int {
	bookCount := map[int]int{}
	for _, book := range books {
		count, ok := bookCount[book]
		if ok {
			bookCount[book] = count + 1
		} else {
			bookCount[book] = 1
		}
	}

	return bookCount
}

func buildDiffCounts(bookCount map[int]int) []int {
	diffCounts := []int{}

	for {
		diffCount := 0
		for key, count := range bookCount {
			if count > 0 {
				bookCount[key] -= 1
				diffCount += 1
			}

			if diffCount == 5 {
				break
			}
		}
		if diffCount > 0 {
			diffCounts = append(diffCounts, diffCount)
		} else {
			break
		}
	}

	return diffCounts
}

func CalBooksPrice(books []int) float64 {
	bookCount := buildBookCount(books)
	diffCounts := buildDiffCounts(bookCount)

	finalPrice := 0.0
	for _, diffCounts := range diffCounts {
		finalPrice += getPriceByDiffCount(diffCounts)
	}

	return finalPrice
}
