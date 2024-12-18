package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type CorruptedMemoryInstructions string

func (cmi *CorruptedMemoryInstructions) read(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	*cmi = CorruptedMemoryInstructions(string(data))
}

func (cmi CorruptedMemoryInstructions) String() string {
	return string(cmi)
}

type Muls [][]int

func (muls *Muls) appendMatch(match []string) {
	var factors []int
	for j := 1; j <= 2; j++ {
		factor, err := strconv.Atoi(match[j])
		if err != nil {
			log.Fatal(err)
		}
		factors = append(factors, factor)
	}
	*muls = append(*muls, []int{factors[0], factors[1]})
}

func (muls *Muls) extract(corruptedMemoryInstructions *CorruptedMemoryInstructions) {
	pattern, err := regexp.Compile(`mul\((?P<f1>[0-9]*),(?P<f2>[0-9]*)\)`)
	if err != nil {
		log.Fatal(err)
	}
	matches := pattern.FindAllStringSubmatch(corruptedMemoryInstructions.String(), -1)
	if matches == nil {
		log.Fatal("No match")
	}
	for _, match := range matches {
		muls.appendMatch(match)
	}
}

func (muls *Muls) extractMoreAccurately(corruptedMemoryInstructions *CorruptedMemoryInstructions) {
	mulExpression, _ := regexp.Compile(`mul\((?P<f1>[0-9]*),(?P<f2>[0-9]*)\)`)
	doExpression, _ := regexp.Compile(`do\(\)`)
	dontExpression, _ := regexp.Compile(`don't\(\)`)
	allMatches := mulExpression.FindAllStringSubmatch(corruptedMemoryInstructions.String(), -1)
	if allMatches == nil {
		log.Fatal("No match")
	}
	mulIndices := mulExpression.FindAllStringIndex(corruptedMemoryInstructions.String(), -1)
	doIndices := doExpression.FindAllStringIndex(corruptedMemoryInstructions.String(), -1)
	dontIndices := dontExpression.FindAllStringIndex(corruptedMemoryInstructions.String(), -1)
	active := true
	j1, j2, j3 := 0, 0, 0
	for i := 0; i < mulIndices[len(mulIndices)-1][1]; {
		if 0 < len(mulIndices) && j1 < len(mulIndices) && i == mulIndices[j1][0] {
			i = mulIndices[j1][1]
			if active {
				muls.appendMatch(allMatches[j1])
			}
			j1 += 1
		} else if 0 < len(doIndices) && j2 < len(doIndices) && i == doIndices[j2][0] {
			i = mulIndices[j2][1]
			j2 += 1
			active = true
		} else if 0 < len(dontIndices) && j3 < len(dontIndices) && i == dontIndices[j3][0] {
			i = dontIndices[j3][1]
			j3 += 1
			active = false
		} else {
			i += 1
		}
	}
}

type SumOfMultiplications int

func (som *SumOfMultiplications) calculate(muls *Muls) {
	sum := 0
	for _, factors := range *muls {
		sum += factors[0] * factors[1]
	}
	*som = SumOfMultiplications(sum)
}

func main() {
	corruptedMemoryInstructions := CorruptedMemoryInstructions("")

	muls := Muls{}
	var sumOfMultiplications SumOfMultiplications
	corruptedMemoryInstructions.read("input.dat")
	muls.extract(&corruptedMemoryInstructions)
	sumOfMultiplications.calculate(&muls)
	fmt.Println("sumOfMultiplications", sumOfMultiplications)

	mulsMoreAccurate := Muls{}
	var sumOfMultiplicationsMoreAccurate SumOfMultiplications
	mulsMoreAccurate.extractMoreAccurately(&corruptedMemoryInstructions)
	sumOfMultiplicationsMoreAccurate.calculate(&mulsMoreAccurate)
	fmt.Println("sumOfMultiplicationsMoreAccurate", sumOfMultiplicationsMoreAccurate)
}
