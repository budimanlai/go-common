package gocommon

import (
	"crypto/rand"
	"strings"
	"unicode"
)

// CapitalizeName mengubah setiap huruf pertama dari kata pada nama menjadi huruf besar,
// dan huruf selanjutnya menjadi huruf kecil.
//
// Contoh penggunaan:
//
//	result := CapitalizeName("jOHN doe")
//	// result == "John Doe"
func CapitalizeName(name string) string {
	words := strings.Fields(name)
	for i, word := range words {
		if len(word) == 0 {
			continue
		}
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		for j := 1; j < len(runes); j++ {
			runes[j] = unicode.ToLower(runes[j])
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

// GenerateRandomString menghasilkan string acak dengan panjang sesuai parameter.
// String terdiri dari angka, huruf besar, huruf kecil, dan simbol '_' serta '-'.
// Jika terjadi error saat pembuatan angka acak, fungsi akan mengembalikan string kosong.
//
// Contoh penggunaan:
//
//	str := GenerateRandomString(12)
//	fmt.Println(str) // Output: "aZ8_9k-L2bQw" (contoh, hasil acak)
func GenerateRandomString(length int) string {
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_-"
	if length <= 0 {
		return ""
	}
	result := make([]byte, length)
	charsetLen := byte(len(charset))
	for i := 0; i < length; i++ {
		b := make([]byte, 1)
		_, err := rand.Read(b)
		if err != nil {
			return ""
		}
		result[i] = charset[int(b[0])%int(charsetLen)]
	}
	return string(result)
}
