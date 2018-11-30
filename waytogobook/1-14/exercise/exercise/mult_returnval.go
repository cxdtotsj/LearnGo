package exercise

func Getnovar(input1 int, input2 int) (int, int, int) {
	sum := input1 + input2
	mult := input1 * input2
	diff := input1 - input2
	return sum, mult, diff
}

func Gethasvar(input1 int, input2 int) (sum, mult, diff int) {
	sum = input1 + input2
	mult = input1 * input2
	diff = input1 - input2
	return
}
