// https://adventofcode.com/2020/day/4

package main

import (
	"2020/input"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

//parse input into array of []passports
func parseInput(input []string) []passport {
	var passportListRaw []string = []string{""}

	//parse individual passport strings into []passportListRaw
	listPosition := 0
	for _, item := range input {
		if item == "" {
			listPosition++
			passportListRaw = append(passportListRaw, "")
			continue
		}
		passportListRaw[listPosition] += item + " "
	}

	//parse passport fields
	passportList := make([]passport, len(passportListRaw))
	for count, item := range passportListRaw {
		fields := strings.Fields(item)
		var pass passport
		for _, field := range fields {
			entry := strings.Split(field, ":")
			switch entry[0] {
			case "byr":
				pass.byr = entry[1]
			case "iyr":
				pass.iyr = entry[1]
			case "eyr":
				pass.eyr = entry[1]
			case "hgt":
				pass.hgt = entry[1]
			case "hcl":
				pass.hcl = entry[1]
			case "ecl":
				pass.ecl = entry[1]
			case "pid":
				pass.pid = entry[1]
			case "cid":
				pass.cid = entry[1]
			}
		}
		passportList[count] = pass
	}

	return passportList
}

// validates that a passport contains required fields
func simpleValidatePassport(p passport) bool {
	//check that required vields are present (except cid)
	if p.byr == "" || p.ecl == "" || p.eyr == "" || p.hcl == "" || p.hgt == "" || p.iyr == "" || p.pid == "" {
		return false
	}
	return true
}

// validates all required field data in passport
func validatePassport(p passport) bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	byr, err := strconv.Atoi(p.byr)
	if err != nil || byr < 1920 || byr > 2002 {
		return false
	}
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	iyr, err := strconv.Atoi(p.iyr)
	if err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	eyr, err := strconv.Atoi(p.eyr)
	if err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}
	// hgt (Height) - a number followed by either cm or in:
	pass, err := regexp.MatchString(`\d+((cm)|(in))`, p.hgt)
	if !pass || err != nil {
		return false
	}
	hgtRegex := regexp.MustCompile(`\d+`)
	hgtVal, _ := strconv.Atoi(hgtRegex.FindString(p.hgt))
	/// If cm, the number must be at least 150 and at most 193.
	cmRegex := regexp.MustCompile(`\d+(cm)`)
	if cmRegex.FindString(p.hgt) != "" {
		if hgtVal < 150 || hgtVal > 193 {
			return false
		}
	}
	/// If in, the number must be at least 59 and at most 76.
	inRegex := regexp.MustCompile(`\d+(in)`)
	if inRegex.FindString(p.hgt) != "" {
		if hgtVal < 59 || hgtVal > 76 {
			return false
		}
	}

	/// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	pass, err = regexp.MatchString(`#[0-9a-f]{6}`, p.hcl)
	if !pass || err != nil {
		return false
	}
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	pass, err = regexp.MatchString(`(amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)`, p.ecl)
	if !pass || err != nil {
		return false
	}
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	pass, err = regexp.MatchString(`\d{9}`, p.pid)
	if !pass || err != nil {
		return false
	}
	// cid (Country ID) - ignored, missing or not.
	return true
}

// validate that passports simply contain required fields
func part1(input []string) int {
	passports := parseInput(input)
	validCount := 0
	for _, passport := range passports {
		if simpleValidatePassport(passport) {
			validCount++
		}
	}

	return validCount
}

// validate passport field data
func part2(input []string) int {
	passports := parseInput(input)
	validCount := 0
	for _, passport := range passports {
		if validatePassport(passport) {
			validCount++
		}
	}

	return validCount - 1
}
