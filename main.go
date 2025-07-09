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

		x := 0

		for index, word := range words {
			re := regexp.MustCompile("[0-9]+")
			matches := re.FindAllString(word, -1)
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

	for i := 1; i < wordCount; i++ {
		x := pos[i+1]
		v := sums[i]

		if i > 1 && pos[i] == 0 {
			continue
		}

		out := fmt.Sprintf("%d", int64(v))
		fmt.Printf("%s%s", out, strings.Repeat(" ", max(x - len(out), 0)))
	}

	if scanner.Err() != nil {
		// Handle error.
	}
}
