package iteration

func Repeat(char string,amount int) (repeated string) {
	for i := 0; i < amount; i++ {
		repeated += char
	}
	return repeated
}
