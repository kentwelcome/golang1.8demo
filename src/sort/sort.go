package main

import (
	"math/rand"
	"sort"
)

type Paper struct {
	Score int
}
type Papers []Paper

func genreateDataArray(len int) []Paper {
	papers := make([]Paper, len)

	for idx := range papers {
		papers[idx].Score = rand.Intn(100)
	}

	return papers
}

func (p Papers) Len() int           { return len(p) }
func (p Papers) Less(i, j int) bool { return p[i].Score < p[j].Score }
func (p Papers) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	papers := genreateDataArray(100000000)
	sort.Sort(Papers(papers))
}
