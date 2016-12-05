package wordbank

import (
	"fmt"
)

func letterDiff(a, b string) int {
	return int(a[0]) - int(b[0])
}

/*
 * The WordBank is a struct that allows quick look-ups
 * for a sortable collection of letters
 */
type WordBank struct {
	indexStore map[string]int // Indicates the current index of the letter
	letters []letterData
}

type letterData struct {
	letter string
	count int
}

func (ld letterData) String() string {
	return fmt.Sprintf("(%v-%v)", ld.letter, ld.count)
}

func NewWordBank() *WordBank {
	return &WordBank{indexStore: make(map[string]int), letters: make([]letterData, 0)}
}

// WordBank interface
func (wb *WordBank) MarkLetter(c string) {
	if idx, found := wb.indexStore[c]; found {
		wb.letters[idx].count++
	} else {
		appendIdx := len(wb.letters)
		wb.letters = append(wb.letters, letterData{c, 1})
		wb.indexStore[c] = appendIdx
	}
}

func (wb *WordBank) LetterIndex(c string) int {
	if idx, found := wb.indexStore[c]; found {
		return idx
	}
	return -1
}

// WordBank sort functions
func (wb *WordBank) Len() int {
	return len(wb.letters)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (wb *WordBank) Less(i, j int) bool {
	letterI, letterJ := wb.letters[i], wb.letters[j]
	countDiff := letterI.count - letterJ.count
	if countDiff == 0 {
		return letterDiff(letterI.letter, letterJ.letter) < 0 // sort alphabetically
	} else {
		return countDiff > 0 // i has more than j, therefore it comes first
	}
}

// Swap swaps the elements with indexes i and j.
func (wb *WordBank) Swap(i, j int) {
	tempI, tempJ := wb.letters[i], wb.letters[j]

	// Swap in the array
	wb.letters[j] = tempI
	wb.letters[i] = tempJ

	// Swap in our tracking map
	wb.indexStore[tempI.letter] = j
	wb.indexStore[tempJ.letter] = i
}

func (wb *WordBank) String() string {
	return fmt.Sprintf("%v", wb.letters)
}
