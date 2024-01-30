package tests

import (
	"fmt"
	"testing"

	"github.com/Ozzy-ZY/proj/internal/iteration"
)

func TestRepeat(t *testing.T) {
	repeated := iteration.Repeat("a", 10)
	exp := "aaaaaaaaaa"
	if repeated != exp {
		t.Errorf("exp %q found %q", exp, repeated)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a", 10)
	}

}
func ExampleRepeat() {
	fmt.Println(iteration.Repeat("a", 5))
	// Output: aaaaa
}
