package main

import (
	"testing"

	util "github.com/MHarland/adventofcode/util"
)

func TestReportsReadFromFile(t *testing.T) {
	reports := Reports{}
	reports.readFromFile("input_test.dat")
	target := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	match := util.MatrixMatch(reports, target)
	if !match {
		t.Errorf("%v != %v", reports, target)
	}
}

func TestReportSafetiesRunChecks(t *testing.T) {
	reports := Reports{}
	safeties := ReportSafeties{}

	reports.readFromFile("input_test.dat")
	safeties.runChecks(&reports)

	target := []bool{true, false, false, false, false, true}
	match := util.SequenceMatch(safeties, target)
	if !match {
		t.Errorf("%v != %v", safeties, target)
	}
}

func TestSafeReportsCount(t *testing.T) {
	reports := Reports{}
	safeties := ReportSafeties{}
	safeReportsCount := SafeReportsCount(0)

	reports.readFromFile("input_test.dat")
	safeties.runChecks(&reports)
	safeReportsCount.count(&safeties)

	target := 2
	match := int(safeReportsCount) == target
	if !match {
		t.Errorf("%v != %v", safeties, target)
	}
}

func TestReportSafetiesRunChecksWithProblemDampener(t *testing.T) {
	reports := Reports{}
	safeties := ReportSafeties{}

	reports.readFromFile("input_test.dat")
	safeties.runChecksWithProblemDampener(&reports)

	target := []bool{true, false, false, true, true, true}
	match := util.SequenceMatch(safeties, target)
	if !match {
		t.Errorf("%v != %v", safeties, target)
	}
}
