package days

import (
	"fmt"
	inputs "github.com/mjbower/adventOfCode/inputs"
)

// https://adventofcode.com/2020/day/6
func Six() {
	inputSlice := inputs.Day6

	var m = map[string]int{}
	totalScore := 0

	for i, s :=  range inputSlice {
		for _, c := range s {
			// set the map value to 1 , to show it exists
			m[string(c)] = 1
		}
		if i > 0 && s == "" {
			// completed group
			totalScore+= len(m)
			fmt.Printf("Group %v scored %d\n",m,len(m))
			// create an empty map for the next group
			m = map[string]int{}
		}
		if i == len(inputSlice)-1 {
			// last line in the file,  so wrap up the group
			totalScore+= len(m)
			fmt.Printf("Group %v scored %d\n",m,len(m))
		}
	}
	fmt.Println("The final score is ", totalScore)
}
