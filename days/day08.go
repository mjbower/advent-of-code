package days

import (
	"fmt"
	inputs "github.com/mjbower/adventOfCode/inputs"
	"os"
	"strconv"
)

// https://adventofcode.com/2020/day/8
var accVal = 0
var seen = map[int]int{}
var ptr = 0
var inputSlice = []string{}

func process (p int,s string) {
	fmt.Printf("OP %s ACC(%d)  PTR(%d)\n",s,accVal,ptr)

	if _, ok := seen[p]; ok {
		fmt.Printf("Already processed, exiting ...   Acc(%d)\n", accVal)
		os.Exit(0)
	}

	myExp := `(?P<op>[a-z]{3})\s(?P<val>[+-][\d+])`
	result := regexToMap(s,myExp)

	newVal, _ := strconv.Atoi(result["val"])

	switch result["op"] {
	case "nop":
			ptr = p + 1
	case "acc":
		accVal += newVal
		ptr = p + 1
	case "jmp":
		ptr = p + newVal
	}

	seen[p] = 1
	process(ptr,inputSlice[ptr])
}

func Eight() {
	inputSlice = inputs.Day8

	// start with the first slice value
	process(0,inputSlice[0])
}