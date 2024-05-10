package vinegar

import (
	"bytes"
	"strings"
)

// VigenereTable produces a Vigenere Table using the given keyword. It first
// will format the keyword given by removing any duplicates and enforcing it
// to be 26 lowercasse latin characters. Once formatted the table is produced
// by shifting the alphabet 26 times maaking a 26x26 matrix of runes.
func VigenereTable(keyword string) [][]rune {
	table := make([][]rune, 26)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	if keyword != "" {
		valid := formatTableKeyword(keyword)
		for _, c := range valid {
			alphabet = strings.ReplaceAll(alphabet, string(c), "")
		}
		alphabet = valid + alphabet
	}

	for i := 0; i < 26; i++ {
		if i > 0 {
			alphabet = alphabet[1:] + string(alphabet[0])
		}
		table[i] = bytes.Runes([]byte(alphabet))
	}

	return table
}

func formatTableKeyword(keyword string) string {
	formatted := strings.ToLower(keyword)
	if containsDuplicate(keyword) {
		formatted = removeDuplicates(keyword)
	}
	if len(formatted) > 26 {
		return formatted[:26]
	}
	return formatted
}

func containsDuplicate(s string) bool {
	seen := make(map[rune]int)
	for _, c := range s {
		seen[c]++
	}
	return len(seen) != len(s)
}

func removeDuplicates(s string) string {
	str := strings.Builder{}
	seen := make(map[rune]int)
	for _, c := range s {
		if _, ok := seen[c]; ok {
			continue
		}
		str.WriteRune(c)
		seen[c]++
	}
	return str.String()
}
