package main
import (
"crypto/rand"
"fmt"
)
func generatePasswordKey() string {
const keyLength = 8
const keyChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
key := make([]byte, keyLength)
_, err := rand.Read(key)
if err != nil {
panic(err)
}
for i := 0; i &lt; keyLength; i++ {
key[i] = keyChars[key[i]%byte(len(keyChars))]
}
return string(key)
}
func main() {
passwordKey := generatePasswordKey()
fmt.Println("Generated password key:", passwordKey)
}