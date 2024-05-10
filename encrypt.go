package vinegar

import (
	"slices"
	"strings"
)

// EncryptVigenere encrypts the given message using the keyword provided
// according to the vigenere table supplied.
func EncryptVigenere(message, keyword string, table [][]rune) string {
	trimmed := strings.ReplaceAll(message, " ", "")
	valid := formatEncryptionKeyword(keyword, trimmed)
	str := strings.Builder{}
	for i := 0; i < len(trimmed); i++ {
		p := rune(trimmed[i])
		k := rune(valid[i])
		idx := slices.Index(table[0], k)
		for _, row := range table {
			if row[0] == p {
				str.WriteRune(row[idx])
			}
		}
	}
	return str.String()
}

// DecryptVigenere decrypts the provided ciphertext using the keyword and
// vigenere table provided - the resulting plaintext will have no spaces.
func DecryptVigenere(cipher, keyword string, table [][]rune) string {
	valid := formatEncryptionKeyword(keyword, cipher)
	str := strings.Builder{}
	for i := 0; i < len(cipher); i++ {
		c := rune(cipher[i])
		k := rune(valid[i])
		idx := slices.Index(table[0], k)
		for _, row := range table {
			if row[idx] == c {
				str.WriteRune(row[0])
			}
		}
	}
	return str.String()
}

func formatEncryptionKeyword(keyword, message string) string {
	k, m := len(keyword), len(message)
	if k == m {
		return keyword
	}
	if k > m {
		return keyword[:m]
	}
	str := strings.Builder{}
	str.WriteString(keyword)
	for str.Len() != m {
		if m-str.Len() >= k {
			str.WriteString(keyword)
		} else {
			str.WriteString(keyword[:m-str.Len()])
		}
	}
	return str.String()
}
