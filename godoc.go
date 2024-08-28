// Package vinegar provides a Vigenere table builder and encryption/decryption
// tools for encoding and decoding plaintext and its corresponding ciphertext.
// Instead of the traditional cipher this package allows for alphabets of
// arbitrary sized UTF-8 strings with spaces and punctuation included.
//
// The Vigenere table requires a TableConfig to be built, this configuration
// defines the alphabet, keyword (for the table) and secret key (for encryption
// and decryption) to be used for the cipher. The table keyword is used to
// shuffle the alphabet by removing all characters from the alphabet that are
// in the keyword (once formatted) and prefixing the alphabet with the keyword.
// This creates a new alphabet to be used as the initial line in the table. If
// the keyword is nil then the alphabet provided will be used as if (once
// formatted). The Secret Key used for en/de-cryption is repeated and truncated
// such that it has the same rune count as the plain/cipher-text and acts as an
// aid for lookups in the Vigenere table. Although supplied in the configuration
// it is not used during en/de-cryption and must be provided instead, in order
// to stop anyone from being able to access table data without the correct key.
package vinegar
