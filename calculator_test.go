package calculator

import "testing"

func assertBooksPrize(t *testing.T, books []int, want float64) {
	t.Helper()

	got := CalBooksPrice(books)

	if got != want {
		t.Errorf("Assert price is %f, but got %f", want, got)
	}
}

func TestCalBooksPrice(t *testing.T) {
	t.Run("Simple cases", func(t *testing.T) {
		examples := []struct {
			name  string
			books []int
			want  float64
		}{
			{"Buy [] book", []int{}, 0.0},
			{"Buy [1] book", []int{1}, price},
			{"Buy [1, 2] book", []int{1, 2}, price * 2 * 0.95},
			{"Buy [1, 2, 3] books", []int{1, 2, 3}, price * 3 * 0.9},
			{"Buy [1, 2, 3, 4] book", []int{1, 2, 3, 4}, price * 4 * 0.8},
			{"Buy [1, 2, 3, 4, 5] book", []int{1, 2, 3, 4, 5}, price * 5 * 0.75},
		}
		for _, example := range examples {
			t.Run(example.name, func(t *testing.T) {
				assertBooksPrize(t, example.books, example.want)
			})
		}
	})

	t.Run("Combo cases", func(t *testing.T) {
		examples := []struct {
			name  string
			books []int
			want  float64
		}{
			{"Buy [1, 1] book", []int{1, 1}, price * 2},
			{"Buy [1, 2, 4] book", []int{1, 2, 4}, price * 3 * 0.9},
			{"Buy [1, 2, 4, 5] book", []int{1, 2, 4, 5}, price * 4 * 0.8},
		}
		for _, example := range examples {
			t.Run(example.name, func(t *testing.T) {
				assertBooksPrize(t, example.books, example.want)
			})
		}
	})
}
