package vinegar

import (
	"errors"
	"unicode/utf8"
)

// TableConfig is the configuration struct for a Vigenere Table which
// defines the Alphabet, Keyword (if any), and SecretKey to be used.
// The TableConfig should only be constructed using NewTableConfig,
// unless proper validations, and standardisations are performed
// on the fields first to conform to the specifications.
type TableConfig struct {
	// Alphabet represents the characters that can be used in the table.
	// It will have any duplicate characters removed as well as any spaces.
	// The alphabet is expected to be a UTF-8 encoded string. It's
	// dimensions are used to generate the table, an alphabet of 26
	// characters (the Latin alphabet for example) will produce a 26x26
	// table.
	//
	// If the keyword is not nil, then the keyword, after being formatted,
	// will be prefixed to the alphabet and its characters removed from
	// the remaining alphabet. For example:
	//	Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" and Keyword = "HELLO"
	//	Alphabet = "HELOABCDFGIJKMNPQRSTUVWXYZ"
	//
	// The alphabet will be shifted in a ring of the alphabet's length to
	// produce a table used for encryption and decryption lookups. A
	// simplified example for what a table would look like is as follows:
	//	ABCDEFGHI
	//	BCDEFGHIA
	//	CDEFGHIAB
	//	DEFGHIABC
	//	EFGHIABCD
	//	FGHIABCDE
	//	GHIABCDEF
	//	HIABCDEFG
	//	IABCDEFGH
	Alphabet []rune
	// Keyword is the secret key used to randomise the table.
	// It will have any duplicate characters removed as well as any
	// characters that are not in the alphabet and spaces. This is
	// used to standardise the keyword such that the table is properly
	// formatted for lookups.
	//
	// If this is set to nil the table will not be altered and will
	// represent the alphabet as provided.
	Keyword []rune
	// SecretKey is the key used for encryption and decryption.
	// It is used to lookup the correct entry in the table according
	// to the plain/cipher text's corresponding index in the table.
	// It will be forced to the same length as the plain/cipher text
	// by either truncating it, or repeating it such that it matches
	// the plain/cipher text's length, for example during encryption:
	//	SecretKey = "HIDDEN", PlainText = "THIS IS A SECRET MESSAGE"
	//	"THISISASECRETMESSAGE"
	//	"HIDDENHIDDENHIDDENHI"
	//
	// Any characters not in the alphabet will be removed as well as
	// any spaces standardising it to match the table. If the Keyword
	// is nil (the table is not shuffled), this will be the only
	// basis for encryption and decryption - it is STRONGLY RECOMMENDED
	// to use both Keyword and SecretKey for a more secure cipher.
	SecretKey []rune
}

func (t *TableConfig) Validate() error {
	if t.Alphabet == nil {
		return errors.New("Nil Alphabet")
	}
	if t.SecretKey == nil {
		return errors.New("Nil Secret Key")
	}
	if !utf8.ValidString(string(t.Alphabet)) {
		return errors.New("Invalid alphabet: not UTF-8 encoded")
	}
	if !utf8.ValidString(string(t.SecretKey)) {
		return errors.New("Invalid secret key: not UTF-8 encoded")
	}
	if t.Keyword != nil {
		if !utf8.ValidString(string(t.Keyword)) {
			return errors.New("Invalid keyword: not UTF-8 encoded")
		}
		if len(t.Keyword) > len(t.Alphabet) {
			return errors.New("Invalid keyword: longer than alphabet")
		}
		// if slices.Contains(t.Keyword, ' ') {
		// 	return errors.New("Invalid keyword: space(s) present")
		// }
		seen := make(map[rune]bool, len(t.Keyword))
		for _, r := range t.Keyword {
			if !seen[r] {
				seen[r] = true
				continue
			}
			return errors.New("Invalid keyword: duplicate character(s)")
		}
	}
	// if slices.Contains(t.Alphabet, ' ') {
	// 	return errors.New("Invalid alphabet: space(s) present")
	// }
	// if slices.Contains(t.SecretKey, ' ') {
	// 	return errors.New("Invalid secret key: space(s) present")
	// }
	seen := make(map[rune]bool, len(t.Alphabet))
	for _, r := range t.Alphabet {
		if !seen[r] {
			seen[r] = true
			continue
		}
		return errors.New("Invalid alphabet: duplicate character(s)")
	}
	return nil
}

// NewTableConfig generates a new table config using the provided alphabet,
// keyword and secretKey - all UTF-8 encoded strings. The alphabet will be
// standardised along with the keyword and secretKey according to their
// individual requirements.
func NewTableConfig(alphabet, keyword, secretKey string) (*TableConfig, error) {
	if alphabet == "" {
		return nil, errors.New("Alphabet cannot be empty")
	}
	if secretKey == "" {
		return nil, errors.New("SecretKey cannot be empty")
	}
	if !utf8.ValidString(alphabet) {
		return nil, errors.New("Invalid alphabet: not UTF-8 encoded")
	}
	if keyword != "" && !utf8.ValidString(keyword) {
		return nil, errors.New("Invalid keyword: not UTF-8 encoded")
	}
	if !utf8.ValidString(secretKey) {
		return nil, errors.New("Invalid secret key: not UTF-8 encoded")
	}

	var stdAlphabet, stdKeyword, stdSecret []rune
	// stdAlphabet = removeSpaces([]rune(alphabet))
	stdAlphabet, _ = removeDuplicates([]rune(alphabet))
	if keyword != "" {
		// stdKeyword = removeSpaces([]rune(keyword))
		stdKeyword, _ = removeDuplicates([]rune(keyword))
		stdKeyword = formatKeyword(stdKeyword, stdAlphabet)
		stdAlphabet = formatAlphabetWithKeyword(stdKeyword, stdAlphabet)
	}
	// stdSecret = removeSpaces([]rune(secretKey))
	stdSecret = formatKeyword([]rune(secretKey), stdAlphabet)
	return &TableConfig{
		Alphabet:  stdAlphabet,
		Keyword:   stdKeyword,
		SecretKey: stdSecret,
	}, nil
}
