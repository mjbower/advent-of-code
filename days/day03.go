package days

import (
	"fmt"
	inputs "github.com/mjbower/adventOfCode/inputs"
)

// https://adventofcode.com/2020/day/3



func Three() {
	inputSlice := inputs.Day3

	treeCount := 0
	skipL, skipD := 3,1
	offset := 0

	for i, s :=  range inputSlice {
		offset = offset + skipL
		fmt.Printf("Line %d offset %d String %s\n",i,offset,s)
		fmt.Println("Char", string(s[offset-1])+ string(s[offset]) + string(s[offset+1]),skipD)
		if string(s[offset]) == "#" {
			treeCount++
		}
	}
	fmt.Printf("You hit %d Trees\n",treeCount)
}
