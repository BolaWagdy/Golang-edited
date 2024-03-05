package tests

import (
	"slices"
	"testing"

	"github.com/Ozzy-ZY/proj/internal/arrays"
)

func TestSum(t *testing.T) {
	t.Run("colletion of 5 items", func(t *testing.T) {
		nums := [5]int{1, 2, 3, 4, 5}
		got := arrays.Sum(nums[0:1]) //the [:] is creating a silce out of
		//the original array and : is a kind of a
		//range operator that is not inclusive
		want := 1
		if got != want {
			t.Errorf("got %d want %d given %d", got, want, nums)
		}
	})
	t.Run("colletion of any size", func(t *testing.T) {
		nums := []int{2, 3, 5}
		got := arrays.Sum(nums[:])
		want := 10
		if want != got {
			t.Errorf("got %d want %d given %d", got, want, nums)
		}
	})
}
func TestSumAll(t *testing.T) {
	got := arrays.SumAll([]int{5, 5}, []int{1, 1, 1})
	want := []int{10, 3}
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
