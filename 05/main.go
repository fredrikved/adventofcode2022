package main

import (
	_ "embed"
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

	instructions := strings.Split(parts[1], "\n")
	r := regexp.MustCompile(`move (?P<Move>\d*) from (?P<From>\d*) to (?P<To>\d*)`)
	log.Default().Print(r.SubexpNames())

	crateList := lo.Values(crates)

	for _, inst := range instructions {
		matches := r.FindStringSubmatch(inst)
		amount, _ := strconv.ParseInt(matches[1], 10, 64)
		from, _ := strconv.ParseInt(matches[2], 10, 64)
		to, _ := strconv.ParseInt(matches[3], 10, 64)

		fromCrate, _ := lo.Find(crateList, func(i *Crate) bool {
			return i.Number == int(from)
		})

		toCrate, _ := lo.Find(crateList, func(i *Crate) bool {
			return i.Number == int(to)
		})

		for i := 0; i < int(amount); i++ {
			item, err := lo.Last(fromCrate.Items)
			if err != nil {
				panic(err)
			}
			fromCrate.Items = fromCrate.Items[:len(fromCrate.Items)-1]

			toCrate.Items = append(toCrate.Items, item)
		}
	}

	str := ""

	for i := 0; i < 9; i++ {
		cr, _ := lo.Find(crateList, func(c *Crate) bool {
			return c.Number == i+1
		})
		ls, _ := lo.Last(cr.Items)
		str += ls
	}

	log.Default().Print("top of boxes: ", str)

	//
	//spew.Dump(crates)
	//spew.Dump(instructions)
}
