package main

import (
	"slices"
	"testing"
)

func TestLettersReadFile(t *testing.T) {
	letters := Letters{}
	letters.readFile("input_test.dat")
	for k, v := range map[[2]int]rune{
		{0, 0}: 'M',
		{0, 3}: 'S',
		{3, 4}: 'A',
	} {
		if letters[k[0]][k[1]] != v {
			t.Errorf("(%v,%v): %v != %v", k[0], k[1], letters[k[0]][k[1]], v)
		}
	}
}

func TestXPositionsScan(t *testing.T) {
	letters := Letters{}
	letters.readFile("input_test.dat")
	xpos := XPositions{}
	xpos.scan(&letters, 'X')
	if !(xpos[0][0] == 0 && xpos[0][1] == 4) {
		t.Errorf("%v, != [0, 4]", xpos[0])
	}
	if !(xpos[len(xpos)-1][0] == 9 && xpos[len(xpos)-1][1] == 9) {
		t.Errorf("%v, != [9, 9]", xpos[len(xpos)-1])
	}
}

func TestXMASMatchesSearchFirstMatch(t *testing.T) {
	letters := Letters{}
	xpos := XPositions{}
	matches := XMASMatches{}
	letters.readFile("input_test.dat")
	xpos.scan(&letters, 'X')
	target := [4][2]int{{0, 5}, {0, 6}, {0, 7}, {0, 8}}
	matches.search(&letters, &xpos)
	if !(slices.Contains(matches, target)) {
		t.Errorf("XMAS not found at [0, 5] - horizontal")
	}
}

func TestXMASMatchesSearchNumberOfMatches(t *testing.T) {
	letters := Letters{}
	xpos := XPositions{}
	matches := XMASMatches{}
	letters.readFile("input_test.dat")
	xpos.scan(&letters, 'X')
	matches.search(&letters, &xpos)
	if !(len(matches) == 18) {
		t.Errorf("%v != 18", len(matches))
	}
}

func TestX_MASMatchesSearchNumberOfMatchesCase02(t *testing.T) {
	letters := Letters{}
	apos := XPositions{}
	matches := X_MASMatches{}
	letters.readFile("input_test02.dat")
	apos.scan(&letters, 'A')
	matches.searchX_Mas(&letters, &apos)
	if !(len(matches) == 1) {
		t.Errorf("%v != 1", len(matches))
	}
}

func TestX_MASMatchesSearchNumberOfMatchesCase03(t *testing.T) {
	letters := Letters{}
	apos := XPositions{}
	matches := X_MASMatches{}
	letters.readFile("input_test03.dat")
	apos.scan(&letters, 'A')
	matches.searchX_Mas(&letters, &apos)
	if !(len(matches) == 9) {
		t.Errorf("%v != 9", len(matches))
	}
}
