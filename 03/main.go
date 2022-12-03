package main

import (
	_ "embed"
	"fmt"
	"github.com/samber/lo"
	"log"
	"strings"
)

//go:embed input.txt
var input string

var characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	sacks := strings.Split(input, "\n")

	totalValue := 0

	for _, sack := range sacks {
		part1 := sack[:len(sack)/2]
		part2 := sack[len(sack)/2:]

		var common string
		for _, char := range part1 {
			str := fmt.Sprintf("%c", char)
			if strings.Index(part2, str) != -1 {
				common = str
			}
		}

		totalValue += strings.Index(characters, common) + 1
	}

	log.Default().Print("P1 - total priority: ", totalValue)

	groups := lo.Reduce(sacks, func(gs [][]string, sack string, index int) [][]string {
		if index%3 == 0 {
			gs = append(gs, []string{sack})
		} else {
			gs[len(gs)-1] = append(gs[len(gs)-1], sack)
		}
		return gs
	}, [][]string{})

	inString := func(source string, char int32) bool {
		for _, ch := range source {
			if ch == char {
				return true
			}
		}
		return false
	}

	totalValue = 0

	for _, group := range groups {
		compare := group[0]
		var common string
		for _, char := range compare {
			if inString(group[1], char) && inString(group[2], char) {
				common = fmt.Sprintf("%c", char)
				break
			}
		}

		totalValue += strings.Index(characters, common) + 1
	}

	log.Default().Print("P2 - total: ", totalValue)
}
