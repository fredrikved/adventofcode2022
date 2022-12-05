package main

import (
	_ "embed"
	"github.com/davecgh/go-spew/spew"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Pair struct {
	SectionA Section
	SectionB Section
}

type Section struct {
	From int
	To   int
}

func (s Section) Contains(sect Section) bool {
	return s.To >= sect.To && s.From <= sect.From
}

func (s Section) Overlaps(sect Section) bool {
	return (s.To >= sect.To && s.From <= sect.To) || (s.From <= sect.From && s.To >= sect.From)
}

func toSection(s string) Section {
	parts := strings.Split(s, "-")
	int1, _ := strconv.ParseInt(parts[0], 10, 64)
	int2, _ := strconv.ParseInt(parts[1], 10, 64)
	return Section{
		From: int(int1),
		To:   int(int2),
	}
}

func main() {
	pairs := lo.Reduce(strings.Split(input, "\n"), func(sr []Pair, s string, _ int) []Pair {
		return append(sr, Pair{
			SectionA: toSection(strings.Split(s, ",")[0]),
			SectionB: toSection(strings.Split(s, ",")[1]),
		})
	}, []Pair{})

	var contains []int

	for index, pair := range pairs {
		if pair.SectionA.Contains(pair.SectionB) || pair.SectionB.Contains(pair.SectionA) {
			contains = append(contains, index)
		}
	}

	spew.Dump("contains: ", len(contains))

	var overlaps []int

	for index, pair := range pairs {
		if pair.SectionA.Overlaps(pair.SectionB) || pair.SectionB.Overlaps(pair.SectionA) {
			overlaps = append(overlaps, index)
		}
	}

	spew.Dump("overlaps: ", len(overlaps))
}
