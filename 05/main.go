package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type PageOrderingRules map[int][]int

func (por *PageOrderingRules) readFromFile(filePath string) {
	var pages []string
	var beforePage, afterPage int
	var err error
	var file []byte

	file, err = os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	firstParagraph := strings.Split(string(file), "\n\n")[0]
	for _, line := range strings.Split(firstParagraph, "\n") {
		pages = strings.Split(line, "|")

		beforePage, err = strconv.Atoi(pages[0])
		if err != nil {
			log.Fatal(err)
		}

		afterPage, err = strconv.Atoi(pages[1])
		if err != nil {
			log.Fatal(err)
		}

		_, ok := (*por)[beforePage]
		if !ok {
			(*por)[beforePage] = []int{}
		}
		(*por)[beforePage] = append((*por)[beforePage], afterPage)
	}
}

type PagesUpdates [][]int

func (pu *PagesUpdates) readFromFile(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	secondParagraph := strings.Split(string(file), "\n\n")[1]
	for _, line := range strings.Split(secondParagraph, "\n") {
		*pu = append(*pu, []int{})
		for _, nr := range strings.Split(line, ",") {
			page, err := strconv.Atoi(nr)
			if err != nil {
				log.Fatal(err)
			}
			(*pu)[len(*pu)-1] = append((*pu)[len(*pu)-1], page)
		}
	}
}

func (pu *PagesUpdates) dropCorrectlyOrdered(orderedPages *PagesUpdatesCorrectlyOrdered) {
	newPu := PagesUpdates{}
	for iOrderedPage, orderedPage := range *orderedPages {
		if !orderedPage {
			newPu = append(newPu, (*pu)[iOrderedPage])
		}
	}
	*pu = newPu
}

func (pu *PagesUpdates) order(rules *PageOrderingRules) {
	for iUpdate := range *pu {
		update := &(*pu)[iUpdate]
		ipage := len(*update) - 1
		for ipage > 0 {
			swapped := false
			for iforbidden := 0; !swapped && iforbidden < len((*rules)[(*update)[ipage]]); iforbidden++ {
				forbidden := (*rules)[(*update)[ipage]][iforbidden]
				for jpage := ipage - 1; !swapped && jpage >= 0; jpage-- {
					if (*update)[jpage] == forbidden {
						tempUpdate := make([]int, 0, len(*update))
						tempUpdate = append(tempUpdate, (*update)[0:jpage]...)
						tempUpdate = append(tempUpdate, (*update)[jpage+1:ipage+1]...)
						tempUpdate = append(tempUpdate, (*update)[jpage])
						tempUpdate = append(tempUpdate, (*update)[ipage+1:]...)
						*update = tempUpdate
						ipage = len(*update)
						swapped = true
					}
				}
			}
			ipage -= 1
		}
	}
}

type PagesUpdatesCorrectlyOrdered []bool

func (puco *PagesUpdatesCorrectlyOrdered) check(rules *PageOrderingRules, updates *PagesUpdates) {
	var isValid bool
	var subSeqToProbe []int
	for _, update := range *updates {
		isValid = true
		for ipage := len(update) - 1; ipage > 0 && isValid; ipage-- {
			subSeqToProbe = update[:ipage]
			for _, forbiddenPage := range (*rules)[update[ipage]] {
				if slices.Contains(subSeqToProbe, forbiddenPage) {
					isValid = false
				}
			}
		}
		*puco = append(*puco, isValid)
	}
}

type SumOfValidUpdateMiddlePageNumbers int

func (sum *SumOfValidUpdateMiddlePageNumbers) sum(updates *PagesUpdates, isOrdered *PagesUpdatesCorrectlyOrdered) {
	var j, s int
	for i := range *updates {
		if (*isOrdered)[i] {
			j = (len((*updates)[i]) - 1) / 2
			s += (*updates)[i][j]
		}
	}
	*sum = SumOfValidUpdateMiddlePageNumbers(s)
}

func (sum *SumOfValidUpdateMiddlePageNumbers) sumAll(updates *PagesUpdates) {
	var j, s int
	for i := range *updates {
		j = (len((*updates)[i]) - 1) / 2
		s += (*updates)[i][j]
	}
	*sum = SumOfValidUpdateMiddlePageNumbers(s)
}

func main() {
	rules := PageOrderingRules{}
	updates := PagesUpdates{}
	ordered := PagesUpdatesCorrectlyOrdered{}
	sum := SumOfValidUpdateMiddlePageNumbers(0)

	rules.readFromFile("input.dat")
	updates.readFromFile("input.dat")
	ordered.check(&rules, &updates)
	sum.sum(&updates, &ordered)
	fmt.Println("sum: ", sum)

	updates.dropCorrectlyOrdered(&ordered)
	updates.order(&rules)
	ordered.check(&rules, &updates)

	sum.sumAll(&updates)
	fmt.Println("sum of corrected updates: ", sum) //5169
}
