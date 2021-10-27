package essentials

func GroupByFrequency(numbers []int) map[int]int {
	frequency := make(map[int]int)
	for _, num := range numbers {
		frequency[num] += 1
	}
	return frequency
}
