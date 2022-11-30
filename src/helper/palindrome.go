package helper

import "fmt"

// ----------------------------------------------------------------
// Check if word is palindrome words
// ----------------------------------------------------------------
// 1. Init class which is contain words data (word).
// 2. Iterate and call checking function.
// 3. Interate word length in reserved arrengements.
// 4. Check if reserved words is match with word
// ----------------------------------------------------------------

func InitializePalindrome() {
	// 1. Init class which is contain words data (word).
	type word = struct {
		word string
	}
	words := make([]word, 0)
	words = append(words, word{word: "racecar"})
	words = append(words, word{word: "bringback"})
	words = append(words, word{word: "neveroddoreven"})
	words = append(words, word{word: "carisonrace"})

	// 2. Iterate and call checking function.
	for _, v := range words {
		/*
			=================
			Palindrome =================
			=================
			true
			false
			true
			false
		*/
		fmt.Println(IsStringReversable(v.word))
	}
}

func IsStringReversable(data string) bool {
	result := []byte{}

	// 3. Interate word length in reserved arrengements.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}

	// 4. Check if reserved words is match with word
	return data == string(result)
}
