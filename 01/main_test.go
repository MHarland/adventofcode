package main

import (
	"testing"

	util "github.com/MHarland/adventofcode/util"
)

func TestInitLocationIds(t *testing.T) {
	locationIds := LocationIds{{}, {}}
	locationIds.initFromFile("input_test.dat")
	target := [][]int{
		{3, 4, 2, 1, 3, 3},
		{4, 3, 5, 3, 9, 3},
	}
	for i, row := range target {
		for j, value := range row {
			if value != locationIds[i][j] {
				t.Errorf("%v != %v", locationIds, target)
			}
		}
	}
}

func TestSortLocationIds(t *testing.T) {
	locationIds := LocationIds{{}, {}}
	locationIds.initFromFile("input_test.dat")
	locationIds.sortLocationIds()
	target := [][]int{
		{1, 2, 3, 3, 3, 4},
		{3, 3, 3, 4, 5, 9},
	}
	match := util.MatrixMatch(locationIds, target)
	if !match {
		t.Errorf("%v != %v", locationIds, target)
	}
}

func TestSubtractPairwise(t *testing.T) {
	locationIds := LocationIds{{}, {}}
	locationIdPairwiseDifferences := LocationIdPairwiseDifferences{}

	locationIds.initFromFile("input_test.dat")
	locationIds.sortLocationIds()

	locationIdPairwiseDifferences.subtractPairwise(&locationIds)
	target := []int{2, 1, 0, 1, 2, 5}
	match := util.SequenceMatch(locationIdPairwiseDifferences, target)
	if !match {
		t.Errorf("%v != %v", locationIdPairwiseDifferences, target)
	}
}

func TestSum(t *testing.T) {
	locationIds := LocationIds{{}, {}}
	locationIdPairwiseDifferences := LocationIdPairwiseDifferences{}
	scalarDifference := ScalarDifference(0)

	locationIds.initFromFile("input_test.dat")
	locationIds.sortLocationIds()

	locationIdPairwiseDifferences.subtractPairwise(&locationIds)

	scalarDifference.sum(locationIdPairwiseDifferences)
	target := 11
	if int(scalarDifference) != target {
		t.Errorf("%v != %v", scalarDifference, target)
	}
}

func TestSimilarityScoreCalculate(t *testing.T) {
	locationIds := LocationIds{{}, {}}
	locationIds.initFromFile("input_test.dat")
	var similarityScore SimilarityScore
	similarityScore.calculate(&locationIds)
	target := 31
	if int(similarityScore) != target {
		t.Errorf("%v != %v", similarityScore, target)
	}
}
