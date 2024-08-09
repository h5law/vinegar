package vinegar

import (
	"slices"
	"strings"
)

// Enforce interface
var _ Vigenere = (*vigenere)(nil)

// vigenere is the implementation of the Vigenere cipher containing the table
// used for encryption and decryption of plain/cipher text as well as the
// configuration used to create the table containing the alphabet and keywords.
type vigenere struct {
	config *TableConfig
	table  [][]rune
}

// NewVigenere produces a Vigenere Table using the given TableConfig.
func NewVigenere(config *TableConfig) (Vigenere, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}
	width := len(config.Alphabet)
	alphabet := make([]rune, width)
	copy(alphabet, config.Alphabet)
	if len(config.Keyword) != 0 {
		alphabet = slices.DeleteFunc(alphabet, func(r rune) bool {
			return slices.Contains(config.Keyword, r)
		})
		alphabet = append(config.Keyword, alphabet...)
	}
	table := make([][]rune, width)
	table[0] = alphabet
	for i := 1; i < width; i++ {
		alphabet = append(alphabet[1:], alphabet[0])
		table[i] = alphabet
	}
	return &vigenere{
		config: config,
		table:  table,
	}, nil
}

// Encrypt encrypts the given message using the keyword provided
// according to the vigenere table of the Vigenere struct.
func (v vigenere) Encrypt(message, keyword string) string {
	runeMsg := []rune(message)
	key := formatSecretKeyword([]rune(keyword), v.config.Alphabet, message)
	str := strings.Builder{}
	for i := 0; i < len(runeMsg); i++ {
		idx := slices.Index(v.table[0][:], key[i])
		for _, row := range v.table {
			if row[0] == runeMsg[i] {
				str.WriteRune(row[idx])
				break
			}
		}
	}
	return str.String()
}

// Decrypt decrypts the provided ciphertext using the keyword and
// vigenere table from the struct
func (v vigenere) Decrypt(cipher, keyword string) string {
	key := formatSecretKeyword([]rune(keyword), v.config.Alphabet, cipher)
	str := strings.Builder{}
	runeCipher := []rune(cipher)
	for i := 0; i < len(runeCipher); i++ {
		idx := slices.Index(v.table[0][:], key[i])
		for _, row := range v.table {
			if row[idx] == runeCipher[i] {
				str.WriteRune(row[0])
				break
			}
		}
	}
	return str.String()
}
