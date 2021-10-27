package essentials

import (
	"fmt"
)

type Dict []string

func (d Dict) Print() {
	fmt.Println("**********************")
	for i, word := range d {
		fmt.Println(i, word)
	}
	fmt.Println("**********************")
}

func (d Dict) Split(threshold int) (Dict, Dict) {
	var bigWords Dict
	var smallWords Dict
	for _, word := range d {
		if len(word) > threshold {
			bigWords = append(bigWords, word)
		} else {
			smallWords = append(smallWords, word)
		}
	}
	return bigWords, smallWords
}
