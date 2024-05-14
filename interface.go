package vinegar

type Vigenere interface {
	// Encrypt encrypts the given message using the keyword provided
	// according to the vigenere table of the Vigenere struct.
	Encrypt(message, keyword string) string
	// Decrypt decrypts the provided ciphertext using the keyword and
	// vigenere table from the struct - the resulting plaintext will have no spaces.
	Decrypt(cipher, keyword string) string
}
