package arrays

func Sum(nums []int) (sum int) {
	sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
