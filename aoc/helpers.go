package aoc

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Contains returns true if the string contains the given substring.
func Contains(str string, substr string) bool {
	for _, i := range substr {
		if !strings.Contains(str, string(i)) {
			return false
		}
	}
	return true
}

// ContainsExactly returns true if the string contains the given substring and has the same length.
func ContainsExactly(str string, substr string) bool {
	if len(str) != len(substr) {
		return false
	}
	for _, i := range substr {
		if !strings.Contains(str, string(i)) {
			return false
		}
	}
	return true
}

// MapSum returns the sum of the values in the map.
func MapSum(theMap map[int]int) int {
	sum := 0
	for _, x := range theMap {
		sum += x
	}
	return sum
}

// Max returns the maximum value in the list.
func Max(list []int) int {
	max := 0
	for _, x := range list {
		if x > max {
			max = x
		}
	}
	return max
}

// ReadFile reads the contents of the file and returns a slice of strings.
func ReadFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	return text
}

// Sum returns the sum of the values in the list.
func Sum(list []int) int {
	sum := 0
	for _, x := range list {
		sum += x
	}
	return sum
}
