package tests
import "testing"
import "github.com/Ozzy-ZY/proj/internal/arrays"
func TestSum(t *testing.T){
	numbers := [5]int{1,2,3,4,5}
	sum := arrays.Sum(numbers[:])
	exp := 15
	if sum != exp {
		t.Errorf("exp '%d' but found '%d' given'%v'",exp,sum,numbers)
	}
}