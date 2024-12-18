package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type LocationIds [][]int

func (lids *LocationIds) initFromFile(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	rows := strings.Split(string(file), "\n")
	for _, row := range rows {
		cols := strings.Split(row, "   ")
		for j, col := range cols {
			v, err := strconv.Atoi(col)
			if err != nil {
				log.Fatal(err)
			}
			(*lids)[j] = append((*lids)[j], v)
		}
	}
}

func (lids *LocationIds) sortLocationIds() {
	for i, _ := range *lids {
		slices.Sort((*lids)[i])
	}
}

type LocationIdPairwiseDifferences []int

func (lpd *LocationIdPairwiseDifferences) subtractPairwise(locationIds *LocationIds) {
	var difference int
	for i, _ := range (*locationIds)[1] {
		if (*locationIds)[0][i] >= (*locationIds)[1][i] {
			difference = (*locationIds)[0][i] - (*locationIds)[1][i]
		} else if (*locationIds)[0][i] < (*locationIds)[1][i] {
			difference = (*locationIds)[1][i] - (*locationIds)[0][i]
		}
		*lpd = append(*lpd, difference)
	}
}

type ScalarDifference int

func (sd *ScalarDifference) sum(summands []int) {
	sum := 0
	for _, summand := range summands {
		sum += summand
	}
	(*sd) = ScalarDifference(sum)
}

type SimilarityScore int

func (ss *SimilarityScore) calculate(lids *LocationIds) {
	score := 0
	for _, vi := range (*lids)[0] {
		for _, vj := range (*lids)[1] {
			if vi == vj {
				score += vi
			}
		}
	}
	*ss = SimilarityScore(score)
}

func main() {
	locationIds := LocationIds{{}, {}}
	locationIdPairwiseDifferences := LocationIdPairwiseDifferences{}
	var scalarDifference ScalarDifference
	var similarityScore SimilarityScore

	locationIds.initFromFile("input.dat")
	locationIds.sortLocationIds()

	locationIdPairwiseDifferences.subtractPairwise(&locationIds)

	scalarDifference.sum(locationIdPairwiseDifferences)
	fmt.Println("scalarDifference", scalarDifference)

	similarityScore.calculate(&locationIds)
	fmt.Println("similarityScore", similarityScore)
}
