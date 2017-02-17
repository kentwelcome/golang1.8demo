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

func main() {
	papers := genreateDataArray(100000000)
	sort.Slice(papers, func(i, j int) bool { return papers[i].Score < papers[j].Score })
}
