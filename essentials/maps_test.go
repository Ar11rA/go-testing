package essentials

import "testing"

func TestGroupByFrequency(t *testing.T) {
	numbers := []int{1, 2, 1, 1, 1, 2, 1, 1, 3, 3, 2}
	frequency := GroupByFrequency(numbers)
	if frequency[1] != 6 {
		t.Errorf("Freq[1] should be 6 %v", frequency[1])
	}
	if frequency[2] != 3 {
		t.Errorf("Freq[2] should be 3 %v", frequency[2])
	}
	if frequency[3] != 2 {
		t.Errorf("Freq[3] should be 2 %v", frequency[3])
	}
}
