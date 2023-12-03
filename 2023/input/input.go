package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// FileStringInput returns a slice of strings for each line in a file
func FileStringInput(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

// FileIntInput returns a slice of ints for each line in a file
func FileIntInput(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			if line == "" {
				continue
			}
			log.Fatal(err)
		}
		input = append(input, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
