package main

import (
	_ "embed"
	"github.com/davecgh/go-spew/spew"
	"github.com/samber/lo"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Crate struct {
	Number int
	Items  []string
}

type Instruction struct {
	Amount int
	From   int
	To     int
}

func main() {
	parts := strings.Split(input, "\n\n")

	crateParts := strings.Split(parts[0], "\n")

	var crates = map[int]*Crate{}

	for _, p := range crateParts {
		if strings.Index(p, "[") == 0 {
			for i, char := range p {
				str := string(char)
				if strings.Index("[] ", str) >= 0 {
					continue
				}
				if _, ok := crates[i]; !ok {
					crates[i] = &Crate{}
				}
				crates[i].Items = append(crates[i].Items, str)
			}
		} else {
			for i, char := range p {
				str := string(char)
				if str != " " {
					num, _ := strconv.ParseInt(str, 10, 64)
					crates[i].Number = int(num)
					crates[i].Items = lo.Reverse(crates[i].Items)
				}
			}
		}
	}

	crateList := lo.Values(crates)

	r := regexp.MustCompile(`move (?P<Move>\d*) from (?P<From>\d*) to (?P<To>\d*)`)
	instructions := lo.Map(strings.Split(parts[1], "\n"), func(inst string, _ int) Instruction {
		matches := r.FindStringSubmatch(inst)
		amount, _ := strconv.ParseInt(matches[1], 10, 64)
		from, _ := strconv.ParseInt(matches[2], 10, 64)
		to, _ := strconv.ParseInt(matches[3], 10, 64)

		return Instruction{
			Amount: int(amount),
			From:   int(from),
			To:     int(to),
		}
	})

	//for _, inst := range instructions {
	//	for i := 0; i < inst.Amount; i++ {
	//		item, err := lo.Last(inst.From.Items)
	//		if err != nil {
	//			panic(err)
	//		}
	//		inst.From.Items = inst.From.Items[:len(inst.From.Items)-1]
	//
	//		inst.To.Items = append(inst.To.Items, item)
	//	}
	//}
	//
	//str := ""
	//
	//for i := 0; i < 9; i++ {
	//	cr, _ := lo.Find(crateList, func(c *Crate) bool {
	//		return c.Number == i+1
	//	})
	//	ls, _ := lo.Last(cr.Items)
	//	str += ls
	//}
	//log.Default().Print("top of boxes: ", str)

	var boxes = lo.Reduce(crateList, func(m map[int][]string, l *Crate, _ int) map[int][]string {
		m[l.Number] = l.Items
		return m
	}, map[int][]string{})

	for _, inst := range instructions {
		from := inst.From
		to := inst.To
		amount := inst.Amount

		items := boxes[from][len(boxes[from])-amount:]

		log.Default().Print(amount)
		log.Default().Print(from, boxes[from])

		boxes[from] = boxes[from][:len(boxes[from])-amount]

		log.Default().Print(to, boxes[to])

		boxes[to] = append(boxes[to], items...)

		sum := 0
		for _, v := range boxes {
			sum += len(v)

		}
		spew.Dump(sum)
	}

	str := ""

	for i := 0; i < 9; i++ {
		items := boxes[i+1]
		ls, _ := lo.Last(items)
		str += ls
	}

	log.Default().Print("top of boxes with multiple: ", str)

	//
	//spew.Dump(crates)
	//spew.Dump(instructions)
}
