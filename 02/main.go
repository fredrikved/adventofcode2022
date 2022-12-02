package main

import (
	_ "embed"
	"github.com/samber/lo"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var scores = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var scoreWin = 6
var scoreDraw = 3

const (
	Rock     = "rock"
	Paper    = "paper"
	Scissors = "scissors"
)

var opMap = map[string]string{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var meMap = map[string]string{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var beats = map[string]string{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

var loses = lo.Invert(beats)

func calculateWins(pairs [][2]string) int {
	totalScore := 0

	for _, pair := range pairs {
		op := pair[0]
		me := pair[1]

		meScore := scores[me]

		score := meScore

		if beats[me] == op {
			score += scoreWin
		} else if me == op {
			score += scoreDraw
		}

		totalScore += score
	}

	return totalScore
}

func main() {
	pairs := lo.Map(strings.Split(input, "\n"), func(i string, _ int) []string {
		return strings.Split(i, " ")
	})

	totalScore := calculateWins(lo.Map(pairs, func(pair []string, _ int) [2]string {
		return [2]string{
			opMap[pair[0]],
			meMap[pair[1]],
		}
	}))

	log.Default().Print("P1 - total score: " + strconv.Itoa(totalScore))

	var newPairs [][2]string

	for _, pair := range pairs {
		op := opMap[pair[0]]
		var me string

		switch pair[1] {
		case "X":
			me = beats[op]
		case "Y":
			me = op
		case "Z":
			me = loses[op]
		}

		newPairs = append(newPairs, [2]string{op, me})
	}

	totalScore = calculateWins(newPairs)

	log.Default().Print("P2 - total score: " + strconv.Itoa(totalScore))
}
