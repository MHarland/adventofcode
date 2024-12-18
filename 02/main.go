package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Reports [][]int

func (r *Reports) readFromFile(filePath string) {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Panic(err)
	}
	fileString := string(fileBytes)
	rows := strings.Split(fileString, "\n")
	for _, row := range rows {
		levels := []int{}
		cols := strings.Split(row, " ")
		for _, col := range cols {
			level, err := strconv.Atoi(col)
			if err != nil {
				log.Panic(err)
			}
			levels = append(levels, level)
		}
		*r = append(*r, levels)
	}
}

type ReportSafeties []bool

func (rs *ReportSafeties) runChecks(reports *Reports) {
	var safeties []bool
	var safe bool
	var allInc, allDec, differ bool
	for _, report := range *reports {
		allInc = checkAllIncreasing(report)
		allDec = checkAllDecreasing(report)
		differ = checkDifferByAtLeastOneAndAtMostThree(report)
		if (allInc || allDec) && differ {
			safe = true
		} else {
			safe = false
		}
		safeties = append(safeties, safe)
	}
	*rs = safeties
}

func (rs *ReportSafeties) runChecksWithProblemDampener(reports *Reports) {
	var reportWithOneLevelOmitted []int
	rs.runChecks(reports)
	for i, safe := range *rs {
		if !safe {
			for j := range (*reports)[i] {
				reportWithOneLevelOmitted = []int{}
				for k, level := range (*reports)[i] {
					if k != j {
						reportWithOneLevelOmitted = append(reportWithOneLevelOmitted, level)
					}
				}
				if (checkAllIncreasing(reportWithOneLevelOmitted) || checkAllDecreasing(reportWithOneLevelOmitted)) && checkDifferByAtLeastOneAndAtMostThree(reportWithOneLevelOmitted) {
					(*rs)[i] = true
					break
				}
			}
		}
	}
}

func checkAllIncreasing(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if !(report[i] < report[i+1]) {
			return false
		}
	}
	return true
}

func checkAllDecreasing(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if !(report[i] > report[i+1]) {
			return false
		}
	}
	return true
}

func checkDifferByAtLeastOneAndAtMostThree(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		difference := report[i] - report[i+1]
		if difference < 0 {
			difference *= -1
		}
		if !(1 <= difference && difference <= 3) {
			return false
		}
	}
	return true
}

type SafeReportsCount int

func (c *SafeReportsCount) count(reportSafeties *ReportSafeties) {
	(*c) = 0
	for _, safe := range *reportSafeties {
		if safe {
			(*c) += 1
		}
	}
}

func main() {
	reports := Reports{}
	safeties := ReportSafeties{}
	safeReportsCount := SafeReportsCount(0)

	reports.readFromFile("input.dat")
	safeties.runChecks(&reports)
	safeReportsCount.count(&safeties)

	fmt.Println("safeReportsCount", safeReportsCount)

	safeties.runChecksWithProblemDampener(&reports)
	safeReportsCount.count(&safeties)
	fmt.Println("safeReportsCount with Problem dampener", safeReportsCount)
}
