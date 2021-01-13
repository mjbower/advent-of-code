package days

import (
	"fmt"
	inputs "github.com/mjbower/adventOfCode/inputs"
	"regexp"
	"strconv"
)

// https://adventofcode.com/2020/day/2

var myExp = regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+)\s(?P<search>[a-z]):\s(?P<password>[a-z]+)`)

func Two() {
	inputSlice := inputs.Day2

	for i, s :=  range inputSlice {
		match := myExp.FindStringSubmatch(s)
		//here we create a named map of the reg results
		result := make(map[string]string)
		for i, name := range myExp.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}

		cRegex := regexp.MustCompile(result["search"])
		matches := cRegex.FindAllStringIndex(result["password"], -1)
		minVal, _ := strconv.Atoi(result["min"])
		maxVal, _ := strconv.Atoi(result["max"])
		if len(matches) >= minVal &&  len(matches) <= maxVal{
			fmt.Printf("Its valid with %d matches!!! -- ",len(matches))
			fmt.Printf("Line %d Min %s Max %s Char %s Password(%s) ---- String %s\n", i,result["min"], result["max"],result["search"],result["password"],s)
		} else {
			fmt.Println("Its Invalid  --",s)
		}
	}
}
