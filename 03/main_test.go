package main

import (
	util "aoc2024/util"
	"testing"
)

func TestCorruptedMemoryInstructionsRead(t *testing.T) {
	memoryInstructions := CorruptedMemoryInstructions("")
	memoryInstructions.read("input_test.dat")

	lenMemoryInstructions := len(memoryInstructions)
	if lenMemoryInstructions < 1 {
		t.Errorf("%v < 1", lenMemoryInstructions)
	}
}

func TestMulsExtract(t *testing.T) {
	corruptedMemoryInstructions := CorruptedMemoryInstructions("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	muls := Muls{}
	target := Muls{{2, 4}, {5, 5}, {11, 8}, {8, 5}}

	muls.extract(&corruptedMemoryInstructions)
	match := util.MatrixMatch(muls, target)
	if !match {
		t.Errorf("%v != %v", muls, target)
	}
}

func TestSumOfMultiplicationsCalculate(t *testing.T) {
	muls := Muls{{2, 4}, {5, 5}, {11, 8}, {8, 5}}
	sumOfMultiplications := SumOfMultiplications(0)
	sumOfMultiplications.calculate(&muls)
	target := SumOfMultiplications(161)
	if sumOfMultiplications != target {
		t.Errorf("%v != %v", sumOfMultiplications, target)
	}
}

func TestMulsExtractMoreAccurately(t *testing.T) {
	corruptedMemoryInstructions := CorruptedMemoryInstructions("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
	muls := Muls{}
	target := Muls{{2, 4}, {8, 5}}

	muls.extractMoreAccurately(&corruptedMemoryInstructions)
	match := util.MatrixMatch(muls, target)
	if !match {
		t.Errorf("%v != %v", muls, target)
	}
}
