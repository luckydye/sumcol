package main

import (
	"bufio"
	"fmt"
	// "log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var sums map[int]float64
	sums = make(map[int]float64)

	for scanner.Scan() {
		str := scanner.Text()

		fmt.Println(str)

		words := strings.Fields(str)

		for index, word := range words {

			re := regexp.MustCompile("[0-9]+")
			matches := re.FindAllString(word, -1)

			if len(matches) > 0 {
				if s, err := strconv.ParseFloat(matches[0], 32); err == nil {
					value := s
					sums[index] = value + sums[index]
				}
			}
		}
	}

	fmt.Println("\nSums:", sums)

	if scanner.Err() != nil {
		// Handle error.
	}
}
