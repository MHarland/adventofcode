package main

import (
	"slices"
	"testing"
)

func TestPageOrderingRulesReadFromFile(t *testing.T) {
	por := PageOrderingRules{}
	por.readFromFile("input_test.dat")
	target := []int{53, 13, 61, 29}
	ok := slices.Equal((por[47]), target)
	if !ok {
		t.Errorf("%v != %v", por[47], target)
	}
}

func TestPagesUpdateReadFromFile(t *testing.T) {
	pu := PagesUpdates{}
	pu.readFromFile("input_test.dat")
	target := []int{75, 47, 61, 53, 29}
	ok := slices.Equal(pu[0], target)
	if !ok {
		t.Errorf("%v != %v", pu[0], target)
	}
}

func TestPagesUpdatesCorrectlyOrderedCheck(t *testing.T) {
	rules := PageOrderingRules{}
	updates := PagesUpdates{}
	ordered := PagesUpdatesCorrectlyOrdered{}
	rules.readFromFile("input_test.dat")
	updates.readFromFile("input_test.dat")
	ordered.check(&rules, &updates)
	target := []bool{true, true, true, false, false, false}
	ok := slices.Equal(ordered, target)
	if !ok {
		t.Errorf("%v != %v", ordered, target)
	}
}

func TestSumSum(t *testing.T) {
	rules := PageOrderingRules{}
	updates := PagesUpdates{}
	ordered := PagesUpdatesCorrectlyOrdered{}
	sum := SumOfValidUpdateMiddlePageNumbers(0)

	rules.readFromFile("input_test.dat")
	updates.readFromFile("input_test.dat")
	ordered.check(&rules, &updates)
	sum.sum(&updates, &ordered)

	target := SumOfValidUpdateMiddlePageNumbers(143)
	ok := sum == target
	if !ok {
		t.Errorf("%v != %v", sum, target)
	}
}

func TestPagesUpdates_dropCorrectlyOrdered(t *testing.T) {
	rules := PageOrderingRules{}
	updates := PagesUpdates{}
	ordered := PagesUpdatesCorrectlyOrdered{}

	rules.readFromFile("input_test.dat")
	updates.readFromFile("input_test.dat")
	ordered.check(&rules, &updates)
	updates.dropCorrectlyOrdered(&ordered)

	targets := [][]int{
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47}}

	for i, target := range targets {
		ok := slices.Equal(updates[i], target)
		if !ok {
			t.Errorf("%v != %v", updates[1], target)
		}
	}
}

func TestPagesUpdates_order(t *testing.T) {
	rules := PageOrderingRules{}
	updates := PagesUpdates{}
	ordered := PagesUpdatesCorrectlyOrdered{}

	rules.readFromFile("input_test.dat")
	updates.readFromFile("input_test.dat")
	ordered.check(&rules, &updates)
	updates.dropCorrectlyOrdered(&ordered)
	updates.order(&rules)

	targets := [][]int{
		{97, 75, 47, 61, 53},
		{61, 29, 13},
		{97, 75, 47, 29, 13}}

	for i, target := range targets {
		ok := slices.Equal(updates[i], target)
		if !ok {
			t.Errorf("%v != %v", updates[1], target)
		}
	}
}
