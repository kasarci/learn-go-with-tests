package arrays

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number;
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		// range keyword lets you iterate over an array.
		// On each iteration, range returns two values - the index and the value. 
		// We are choosing to ignore the index value by using _ blank identifier.
		sums = append(sums, Sum(numbers))
	}
	return sums
}