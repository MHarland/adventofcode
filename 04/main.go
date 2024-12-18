package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Letters [][]rune

func (l *Letters) readFile(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	iRow := len(*l)
	*l = append(*l, []rune{})
	for _, char := range string(file) {
		if char == '\n' {
			iRow += 1
			*l = append(*l, []rune{})
		} else {
			(*l)[iRow] = append((*l)[iRow], rune(char))
		}
	}
}

func (l *Letters) getSequenceOfSize4(path *[4][2]int) string {
	seq := ""
	char := ""
	for i := range path {
		xi := (*path)[i][0]
		if 0 <= xi && xi < len(*l) {
			yi := (*path)[i][1]
			if 0 <= yi && yi < len((*l)[xi]) {
				char = string((*l)[xi][yi])
				seq = strings.Join([]string{seq, char}, "")
			} else {
				return ""
			}
		} else {
			return ""
		}
	}
	return seq
}

func (l *Letters) pickXTopLeftTopRightBotRightBotLeft(center [2]int) string {
	seq := ""
	char := ""
	coords := [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
	var xi, yi int
	for i := range coords {
		xi = center[0] + coords[i][0]
		if 0 <= xi && xi < len(*l) {
			yi = center[1] + coords[i][1]
			if 0 <= yi && yi < len((*l)[xi]) {
				char = string((*l)[xi][yi])
				seq = strings.Join([]string{seq, char}, "")
			} else {
				return ""
			}
		} else {
			return ""
		}
	}
	return seq
}

type XPositions [][2]int

func (xp *XPositions) scan(letters *Letters, x rune) {
	for i := range *letters {
		for j := range (*letters)[i] {
			if (*letters)[i][j] == x {
				*xp = append(*xp, [2]int{i, j})
			}
		}
	}
}

func scanRight(x [2]int, path *[4][2]int) {
	for i := range *path {
		(*path)[i][0] = x[0]
		(*path)[i][1] = x[1] + i
	}
}

func scanLeft(x [2]int, path *[4][2]int) {
	for i := range *path {
		(*path)[i][0] = x[0]
		(*path)[i][1] = x[1] - i
	}
}

func scanDown(x [2]int, path *[4][2]int) {
	for i := range *path {
		(*path)[i][0] = x[0] + i
		(*path)[i][1] = x[1]
	}
}

func scanUp(x [2]int, path *[4][2]int) {
	for i := range *path {
		(*path)[i][0] = x[0] - i
		(*path)[i][1] = x[1]
	}
}

func scanDownRight(x [2]int, path *[4][2]int) {
	for i := range *path {
		(*path)[i][0] = x[0] + i
		(*path)[i][1] = x[1] + i
	}
}

func scanDownLeft(x [2]int, path *[4][2]int) {
	for i := range *path {
		(*path)[i][0] = x[0] + i
		(*path)[i][1] = x[1] - i
	}
}

func scanUpRight(x [2]int, path *[4][2]int) {
	for i := range *path {
		(*path)[i][0] = x[0] - i
		(*path)[i][1] = x[1] + i
	}
}

func scanUpLeft(x [2]int, path *[4][2]int) {
	for i := range *path {
		(*path)[i][0] = x[0] - i
		(*path)[i][1] = x[1] - i
	}
}

type XMASMatches [][4][2]int

func (xm *XMASMatches) search(letters *Letters, xpos *XPositions) {
	var scanPath [4][2]int
	var matchAttempt string
	pathGenerators := []func([2]int, *[4][2]int){
		scanRight,
		scanLeft,
		scanUp,
		scanDown,
		scanUpRight,
		scanUpLeft,
		scanDownLeft,
		scanDownRight,
	}

	for ixpos := range *xpos {
		for _, pathGenerator := range pathGenerators {
			pathGenerator((*xpos)[ixpos], &scanPath)
			matchAttempt = letters.getSequenceOfSize4(&scanPath)
			if matchAttempt == "XMAS" {
				*xm = append(*xm, scanPath)
			}
		}
	}
}

type X_MASMatch struct {
	center    [2]int
	signature string // TopLeftTopRightBotRightBotLeft
}

type X_MASMatches []X_MASMatch

func (xm *X_MASMatches) searchX_Mas(letters *Letters, apos *XPositions) {
	var signature string
	targetSignatures := []string{
		"MMSS", "SMMS", "SSMM", "MSSM",
	}
	for _, apo := range *apos {
		signature = letters.pickXTopLeftTopRightBotRightBotLeft(apo)
		if slices.Contains(targetSignatures, signature) {
			match := &X_MASMatch{
				center:    apo,
				signature: signature,
			}
			*xm = append(*xm, *match)
		}
	}
}

func main() {
	letters := Letters{}
	xpos := XPositions{}
	apos := XPositions{}
	matches := XMASMatches{}
	x_masMatches := X_MASMatches{}

	letters.readFile("input.dat")
	xpos.scan(&letters, 'X')
	matches.search(&letters, &xpos)
	fmt.Println("Number of XMASMatches:", len(matches))

	apos.scan(&letters, 'A')
	x_masMatches.searchX_Mas(&letters, &apos)
	fmt.Println("Number of X-MASMatches:", len(x_masMatches))
}
