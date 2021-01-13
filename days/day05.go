package days

import (
	"fmt"
	inputs "github.com/mjbower/adventOfCode/inputs"
	"regexp"
	"strings"
)

// https://adventofcode.com/2020/day/5

//takes a string and a regex string, compiles it returns a map[string][string] of named parameters
func regexToMap(searchString string, r string ) map[string]string {
	//fmt.Printf("XError: Regex did not match : %q %v\n", searchString,r)
	var regEx = regexp.MustCompile(r)
	n1 := regEx.SubexpNames()
	r2s := regEx.FindAllStringSubmatch(searchString, -1)
	if len(r2s) == 0 {
		fmt.Printf("Error: Regex did not match : %q %v\n", searchString,regEx)
		return nil
	}
	r2 := r2s[0]

	md := map[string]string{}
	for i, n := range r2 {
		md[n1[i]] = n
	}
	//fmt.Printf("XMatch %v\n",md)
	return md
}

func rowCalc(row string) int{
	s := strings.Split(row, "")
	idx := 0
	if s[0] == "B" { idx = 64 }
	if s[1] == "B" { idx += 32 }
	if s[2] == "B" { idx += 16 }
	if s[3] == "B" { idx += 8 }
	if s[4] == "B" { idx += 4 }
	if s[5] == "B" { idx += 2 }
	if s[6] == "B" { idx += 1 }

	return idx
}

func seatCalc(seat string) int {
	s := strings.Split(seat, "")
	// start at 0,  move the index to the right depending on the value of R or L
	idx := 0
	if s[0] == "R" { idx = 4 }
	if s[1] == "R" { idx += 2 }
	if s[2] == "R" { idx += 1 }
	return idx
}

func Five() {
	inputSlice := inputs.Day5
	// BFFFBBFRRR   7 character row identifier, 3 character seat number
	var bPassExp = `(?P<row>[A-Z]{7})(?P<seat>[A-Z]{3})`
	highestSeatID := 0
	for i, s :=  range inputSlice {
		fmt.Printf("Line %d String %s  ",i,s)
		pPort := regexToMap(s,bPassExp)
		row := rowCalc(pPort["row"])
		seat := seatCalc(pPort["seat"])
		seatID := row * 8 + seat
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
		fmt.Printf("Match %v Row(%d) Seat (%d) SeatID (%d)\n",pPort, row,seat,seatID)
	}
	fmt.Println("Highest seatID = ",highestSeatID)
}
