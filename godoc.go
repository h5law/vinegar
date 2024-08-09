// Package vinegar provides a Vigenere table builder and encryption/decryption
// tools for encoding and decoding plaintext and its corresponding ciphertext.
// Instead of the traditional cipher this package allows for alphabets of
// arbitrary sized UTF-8 strings with spaces and punctuation included.
//
// The Vigenere table requires a TableConfig to be built, this configuration
// defines the alphabet, keyword (for the table) and secret key (for encryption
// and decryption) to be used for the cipher.
package vinegar
