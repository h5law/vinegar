package vinegar

import (
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
	v.table[0] = ([26]rune)([]rune(alphabet))
	for i := 1; i < 26; i++ {
		alphabet = alphabet[1:] + string(alphabet[0])
		v.table[i] = ([26]rune)([]rune(alphabet))
	}
	return v
}

// Encrypt encrypts the given message using the keyword provided
// according to the vigenere table of the Vigenere struct.
func (v *vigenere) Encrypt(message, keyword string) string {
	msg := formatKeyword(message)
	key := formatEncryptionKeyword(keyword, msg)
	str := strings.Builder{}
	for i := 0; i < len(msg); i++ {
		p := rune(msg[i])
		idx := slices.Index(v.table[0][:], rune(key[i]))
		for _, row := range v.table {
			if row[0] == p {
				str.WriteRune(row[idx])
				break
			}
		}
	}
	return str.String()
}

// Decrypt decrypts the provided ciphertext using the keyword and
// vigenere table from the struct - the resulting plaintext will have no spaces.
func (v *vigenere) Decrypt(cipher, keyword string) string {
	key := formatEncryptionKeyword(keyword, cipher)
	str := strings.Builder{}
	for i := 0; i < len(cipher); i++ {
		c := rune(cipher[i])
		idx := slices.Index(v.table[0][:], rune(key[i]))
		for _, row := range v.table {
			if row[idx] == c {
				str.WriteRune(row[0])
				break
			}
		}
	}
	return str.String()
}
