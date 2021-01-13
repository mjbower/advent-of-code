package days

import (
	"fmt"
	inputs "github.com/mjbower/adventOfCode/inputs"
	"strconv"
)

// https://adventofcode.com/2020/day/9

pwd

func checkVals(*v []int, i int) {

}

func Nine() {
	inputSlice = inputs.Day9
	values := make([]int, len(inputSlice))

	//first lets get out file into a slice of numbers
	for i, s :=  range inputSlice {
		v,_ := strconv.Atoi(s)
		values[i] = v
		fmt.Printf("Line (%d) Val(%s)\n",i,s)
	}
	fmt.Println(values)

	checkVals(&values, 5)
}