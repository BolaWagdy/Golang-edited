package arrays

func Sum(nums []int) (sum int) {
	sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func SumAll(numtosum ...[]int) (sums []int) {
	for _, nums := range numtosum {
		sums = append(sums, Sum(nums))
	}
	return sums
}
