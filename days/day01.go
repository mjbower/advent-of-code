package days

import (
	"fmt"
	inputs "github.com/mjbower/adventOfCode/inputs"
	"os"
)

// https://adventofcode.com/2020/day/1

func One() {
	inputSlice := inputs.Day1

	for i := 0; i < len(inputSlice); i++ {
		for j := 0; j < len(inputSlice); j++ {
			if i == j { break }
			if (inputSlice[i] + inputSlice[j] ) == 2020 {
				fmt.Printf("Found %d row(%d) %d row(%d) Total = %d\n",inputSlice[i],i,inputSlice[j],j,inputSlice[i] * inputSlice[j])
				os.Exit(0)
			}
		}
	}
}