package vinegar

import (
	"bytes"
	"slices"
	"strings"
)

// Enforce interface
var _ Vigenere = (*vigenere)(nil)

// vigenere is the implementation of the Vigenere cipher containing the table
// used for encryption and decryption of plain/cipher text
type vigenere struct {
	table [26][26]rune
}

// NewVigenere produces a Vigenere Table using the given keyword. It first
// will format the keyword given by removing any duplicates and enforcing it
// to be 26 lowercasse latin characters. Once formatted the table is produced
// by shifting the alphabet 26 times maaking a 26x26 matrix of runes
func NewVigenere(keyword string) Vigenere {
	v := &vigenere{
		table: [26][26]rune{},
	}
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
		copy(v.table[i][:], bytes.Runes([]byte(alphabet)))
	}
	return v
}

// Encrypt encrypts the given message using the keyword provided
// according to the vigenere table of the Vigenere struct.
func (v *vigenere) Encrypt(message, keyword string) string {
	trimmed := strings.ReplaceAll(message, " ", "")
	valid := formatEncryptionKeyword(keyword, trimmed)
	str := strings.Builder{}
	for i := 0; i < len(trimmed); i++ {
		p := rune(trimmed[i])
		k := rune(valid[i])
		idx := slices.Index(v.table[0][:], k)
		for _, row := range v.table {
			if row[0] == p {
				str.WriteRune(row[idx])
			}
		}
	}
	return str.String()
}

// Decrypt decrypts the provided ciphertext using the keyword and
// vigenere table from the struct - the resulting plaintext will have no spaces.
func (v *vigenere) Decrypt(cipher, keyword string) string {
	valid := formatEncryptionKeyword(keyword, cipher)
	str := strings.Builder{}
	for i := 0; i < len(cipher); i++ {
		c := rune(cipher[i])
		k := rune(valid[i])
		idx := slices.Index(v.table[0][:], k)
		for _, row := range v.table {
			if row[idx] == c {
				str.WriteRune(row[0])
			}
		}
	}
	return str.String()
}
