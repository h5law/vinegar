package vinegar

import (
	"slices"
	"unicode/utf8"
)

// removeSpaces removes all the spaces in a given rune slice.
// func removeSpaces(r []rune) []rune {
// 	if !utf8.ValidString(string(r)) {
// 		panic("Invalid UTF-8 string") // this should never happen
// 	}
// 	for i := slices.Index(r, ' '); i != -1; i = slices.Index(r, ' ') {
// 		r = append(r[:i], r[i+1:]...)
// 	}
// 	return r
// }

// removeDuplicateS removes all duplicates from
// a UTF-8 string returning a slice of runes and its length.
func removeDuplicates(word []rune) ([]rune, int) {
	width := len(word) // length in runes not bytes
	seen := make(map[rune]bool, width)
	stdWord := make([]rune, 0, width)

	for len(word) > 0 {
		if !seen[word[0]] {
			stdWord = append(stdWord, word[0])
			seen[word[0]] = true
		}
		word = word[1:]
	}

	return stdWord, len(stdWord) // length in runes not bytes
}

// formatKeyword removes any characters from the keyword that are not in
// the alphabet preparing it to be prefixed to the alphabet for the table.
func formatKeyword(keyword, alphabet []rune) []rune {
	alphabetMap := make(map[rune]bool)
	for _, r := range alphabet {
		alphabetMap[r] = true
	}
	stdKeyword := make([]rune, 0, len(keyword))
	for len(keyword) > 0 {
		if alphabetMap[keyword[0]] {
			stdKeyword = append(stdKeyword, keyword[0])
		}
		keyword = keyword[1:]
	}
	if len(stdKeyword) > len(alphabet) {
		stdKeyword = stdKeyword[:len(alphabet)]
	}
	return stdKeyword
}

// formatAlphabetWithKeyword shuffles the provided alphabet removing any
// characters that are in the keyword and prefixing it with the keyword.
func formatAlphabetWithKeyword(keyword, alphabet []rune) []rune {
	if len(keyword) != 0 {
		// ensure the keyword is formatted correctly
		keyword = formatKeyword(keyword, alphabet)
		// remove any characters in the alphabet that are in the keyword
		alphabet = slices.DeleteFunc(alphabet, func(r rune) bool {
			return slices.Contains(keyword, r)
		})
		// prefix the alphabet with the keyword
		alphabet = append(keyword, alphabet...)
	}
	return alphabet
}

// formatEncryptionKeyword ensures the keyword is formatted with formatKeyword
// then is repeated (and spliced if necessary) to be the same length as the message.
func formatSecretKeyword(secret, alphabet []rune, message string) []rune {
	if !utf8.ValidString(message) {
		panic("Invalid UTF-8 string") // this should never happen
	}
	secret = formatKeyword(secret, alphabet) // ensure secret is formatted correctly

	runeMsg := []rune(message)
	// runeMsg = removeSpaces(runeMsg) // spaces will form a pattern as they are so common

	k, m := len(secret), len(runeMsg)
	if k == m {
		return secret
	}
	// if k > m {
	// 	return secret[:m]
	// }
	paddedSecret := make([]rune, 0, len(runeMsg))
	for len(paddedSecret) != m {
		if m-len(paddedSecret) >= k {
			paddedSecret = append(paddedSecret, secret...)
		} else {
			paddedSecret = append(paddedSecret, secret[:m-len(paddedSecret)]...)
		}
	}

	return paddedSecret
}
