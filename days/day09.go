package days

import (
	"fmt"
	inputs "github.com/mjbower/adventOfCode/inputs"
	"os"
	"strconv"
)

// https://adventofcode.com/2020/day/9

func Nine() {
	inputSlice = inputs.Day9
	values := make([]int, len(inputSlice))

	//first lets get out file into a slice of numbers
	for i, s :=  range inputSlice {
		v,_ := strconv.Atoi(s)
		values[i] = v
	}

	ptr := 5
	for i, _ :=  range values {
		if i < ptr {
			continue
		}

		checkVal := values[ptr]
		fmt.Printf("Checking for (%d) - ",checkVal)
		found := false
		for outer := ptr-5; outer < ptr; outer++ {
			for inner := outer+1; inner < ptr; inner++ {
				if values[outer] + values[inner] == checkVal {
					found = true
					fmt.Printf("  Found %d + %d = %d\n",values[outer],values[inner],checkVal)
					break
				}
			}
		}
		if ! found {
			fmt.Println("Game over, not found")
			os.Exit(0)
		}
		ptr++
	}
}