package main

import (
	"2020/input"
	"testing"
)

var entries []string = []string{
	"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
	"byr:1937 iyr:2017 cid:147 hgt:183cm",
	"",
	"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
	"hcl:#cfa07d byr:1929",
	"",
	"hcl:#ae17e1 iyr:2013",
	"eyr:2024",
	"ecl:brn pid:760753108 byr:1931",
	"hgt:179cm",
	"",
	"hcl:#cfa07d eyr:2025 pid:166559648",
	"iyr:2011 ecl:brn hgt:59in",
}

func TestPart1(t *testing.T) {
	expected := 2

	result := part1(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 2

	result := part2(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2Valid(t *testing.T) {
	input := input.GetStringInput("test-valid.txt")
	expected := 4

	result := part2(input)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2Invalid(t *testing.T) {
	input := input.GetStringInput("test-invalid.txt")
	expected := 0

	result := part2(input)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
