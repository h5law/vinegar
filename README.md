# vinegar

vinegar is a library for producing, Vigenere tables and encrypting/decrypting
messages with them.

Plaintext messages have their spaces removed during encryption as the library
uses the standard latin alphabet modified by a keyword (with duplicates removed).
This means the decrypted message will also have no spaces but will match the
plaintext input in all other ways.
