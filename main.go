package main

import (
	"bufio"
	"fmt"

	// "log"
	"os"
	"regexp"
	"strconv"
	"strings"
	// "math"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sums := make(map[int]float64)
	pos := make(map[int]int)
	wordCount := 0

	for scanner.Scan() {
		str := scanner.Text()

		fmt.Println(str)

		words := strings.Fields(str)
		// wordsRaw := strings.Split(str, " ")

		x := 0

		for index, word := range words {
			matches := regexp.MustCompile("[0-9]+").FindAllString(word, -1)

			wordCount += 1
			x += 1

			if len(matches) > 0 {
				if s, err := strconv.ParseFloat(matches[0], 32); err == nil {
					sums[index] = s + sums[index]
					pos[index] = x
					x = 0
				}
			}

			x += len(word)
		}
	}

	fmt.Println("")
	fmt.Println("Total:")

	fmt.Printf("%s", strings.Repeat(" ", 14))

	for i := 1; i < wordCount; i++ {
		x := pos[i+1]
		v := sums[i]

		out := fmt.Sprintf("%d", int64(v))

		if v == 0 {
			out = ""
		}

		fmt.Printf("%s%s", out, strings.Repeat(" ", max(x - len(out), 0)))
	}

	fmt.Println("")

	if scanner.Err() != nil {
		// Handle error.
	}
}
