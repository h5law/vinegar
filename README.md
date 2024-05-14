# vinegar

vinegar is a library for producing, Vigenere tables and encrypting/decrypting
messages with them.

Plaintext messages have their spaces removed during encryption as the library
uses the standard latin alphabet modified by a keyword (with duplicates removed).
This means the decrypted message will also have no spaces but will match the
plaintext input in all other ways.

```go
package main

import (
    "fmt"
  
    "github.com/h5law/vinegar"
)

func main() {
    vig := vinegar.NewVigenere("kyrptos")
    cipher := vig.Encrypt("this is a secret hidden message", "neddih")
    plain := vig.Decrypt(cipher, "neddih")
    fmt.Println(cipher)                        // wzzjtqklunlzwzzqzzfpujuusv
    fmt.Println(plain)                         // thisisasecrethiddenmessage
    // Using the wrong key
    fmt.Println(vig.Decrypt(cipher, "hidden")) // cdisnyfrecsscdidhsxhesdrma
}
```
