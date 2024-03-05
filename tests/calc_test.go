package tests

import (
	"testing"

	"github.com/Ozzy-ZY/proj/internal/calculator"
)

func TestCalc(t *testing.T) {
	t.Run("addition", func(t *testing.T) {
		got := calculator.Calc(3, '+', 2)
		want := 5
		if want != got {
			t.Errorf("got: %d want %d", got, want)
		}
	})
	t.Run("subtration", func(t *testing.T) {
		got := calculator.Calc(3, '-', 2)
		want := 1
		if want != got {
			t.Errorf("got: %d want %d", got, want)
		}
	})
	t.Run("multiply", func(t *testing.T) {
		got := calculator.Calc(3, '*', 2)
		want := 6
		if want != got {
			t.Errorf("got: %d want %d", got, want)
		}
	})
	t.Run("Division", func(t *testing.T) {
		got := calculator.Calc(3, '/', 2)
		want := 1
		if want != got {
			t.Errorf("got: %d want %d", got, want)
		}
	})
}
