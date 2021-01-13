package days

import (
	"fmt"
	inputs "github.com/mjbower/adventOfCode/inputs"
	"strings"
)

// https://adventofcode.com/2020/day/4

func checkPassport(p map[string]string) bool {
	elem := len(p)

	fmt.Printf("New passport %v ELEM %d",p,elem)
	if elem == 8 {
		fmt.Println(" Good")
		return true
	}
	if elem ==7 {
		if _, ok := p["cid"]; ok {
			fmt.Println(" Bad")
			return false
		} else {
			fmt.Println(" Good")
			return true
		}
	}
	fmt.Println(" Bad")
	return false
}

func Four() {
	inputSlice := inputs.Day4

	var m map[string]string
	var ss []string
	validPassports := 0
	m = make(map[string]string)

	for i, s :=  range inputSlice {
		//fmt.Printf("Line %d String %s\n",i,s)
		if s == "" {
			if checkPassport(m) {
				validPassports++
			}
			m = make(map[string]string)
			continue
		}
		ss = strings.Split(s, " ")

		for _, pair := range ss {
			z := strings.Split(pair, ":")
			m[z[0]] = z[1]
		}
		// capture last passport
		if i == len(inputSlice)-1 {
			if checkPassport(m) {
				validPassports++
			}
			m = make(map[string]string)
			continue
		}
	}
	fmt.Println("Valid Passports = ",validPassports)

}
