package gocommon

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// PtrToFloat64 mengembalikan nilai float64 dari pointer, atau 0 jika nil.
func PtrToFloat64(ptr *float64) float64 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// NormalizePhoneNumber menormalkan string nomor telepon Indonesia ke format internasional yang diawali dengan "62".
// Fungsi ini menghapus semua karakter non-digit dari input, lalu menerapkan aturan berikut:
//   - Jika nomor diawali dengan "62", maka dikembalikan apa adanya.
//   - Jika nomor diawali dengan "0", maka "0" diganti dengan "62".
//   - Jika nomor diawali dengan "8", maka "62" ditambahkan di depan nomor.
//   - Selain itu, string digit yang sudah dibersihkan dikembalikan tanpa perubahan.
//
// Contoh:
//
//	NormalizePhoneNumber("+62 812-3456-7890") // mengembalikan "6281234567890"
//	NormalizePhoneNumber("081234567890")      // mengembalikan "6281234567890"
//	NormalizePhoneNumber("81234567890")       // mengembalikan "6281234567890"
func NormalizePhoneNumber(phone string) string {
	// Hapus semua karakter non-digit
	var digits strings.Builder
	for _, r := range phone {
		if unicode.IsDigit(r) {
			digits.WriteRune(r)
		}
	}
	num := digits.String()

	// Normalisasi prefix
	if strings.HasPrefix(num, "62") {
		return num
	}
	if strings.HasPrefix(num, "0") {
		return "62" + num[1:]
	}
	if strings.HasPrefix(num, "8") {
		return "62" + num
	}
	return num
}

// GenerateTransactionID menghasilkan string ID transaksi unik.
// ID terdiri dari timestamp dengan format "yymmddHHMMSS" diikuti
// 8 karakter heksadesimal acak, memastikan keunikan dan keterlacakan.
// Fungsi ini akan panic jika gagal menghasilkan byte acak.
//
// Contoh penggunaan:
//
//	id := GenerateTransactionID()
//	fmt.Println(id) // Output: "24060712345612345678" (contoh, hasil acak)
func GenerateTransactionID() string {
	now := time.Now()
	timestamp := now.Format("060102150405") // "06" untuk 2 digit tahun

	// Generate 8 random digits
	randomDigits := make([]byte, 8)
	for i := 0; i < 8; i++ {
		b := make([]byte, 1)
		_, err := rand.Read(b)
		if err != nil {
			// Return a fallback value or empty string if random generation fails
			return ""
		}
		randomDigits[i] = '0' + (b[0] % 10)
	}

	return fmt.Sprintf("%s%s", timestamp, string(randomDigits))
}

// GenerateUnique6Digits menghasilkan dan mengembalikan string yang terdiri dari 6 digit angka acak.
// Fungsi ini menggunakan random kriptografi untuk memastikan keunikan dan tidak mudah ditebak.
// Jika terjadi error saat pembuatan angka acak, fungsi akan mengembalikan string kosong.
//
// Contoh penggunaan:
//
//	code := GenerateUnique6Digits()
//	fmt.Println(code) // Output: "123456" (contoh, hasil acak)
func GenerateUnique6Digits() string {
	digits := make([]byte, 6)
	for i := 0; i < 6; i++ {
		b := make([]byte, 1)
		_, err := rand.Read(b)
		if err != nil {
			return ""
		}
		digits[i] = '0' + (b[0] % 10)
	}
	return string(digits)
}

// GenerateUUIDv4 menghasilkan UUID versi 4 (random).
// UUID v4 terdiri dari 16 byte acak dengan beberapa bit diatur sesuai standar RFC 4122.
// Fungsi ini mengembalikan string UUID dalam format standar, atau string kosong jika terjadi error.
//
// Contoh penggunaan:
//
//	uuid := GenerateUUIDv4()
//	fmt.Println(uuid) // Output: "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx"
func GenerateUUIDv4() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return ""
	}

	// Set versi (4) dan varian (RFC 4122)
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10xxxxxx

	// Format ke string UUID standar (36 karakter)
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%04x%08x",
		uint32(uuid[0])<<24|uint32(uuid[1])<<16|uint32(uuid[2])<<8|uint32(uuid[3]),
		uint16(uuid[4])<<8|uint16(uuid[5]),
		uint16(uuid[6])<<8|uint16(uuid[7]),
		uint16(uuid[8])<<8|uint16(uuid[9]),
		uint16(uuid[10])<<8|uint16(uuid[11]),
		uint32(uuid[12])<<24|uint32(uuid[13])<<16|uint32(uuid[14])<<8|uint32(uuid[15]),
	)
}

// HashPassword menghasilkan hash bcrypt dari password plaintext.
// Fungsi ini mengembalikan hash sebagai string, atau string kosong jika terjadi error.
//
// Contoh penggunaan:
//
//	hash := HashPassword("passwordku")
//	fmt.Println(hash) // Output: "$2a$10$..."
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

// CheckPasswordHash memvalidasi apakah password plaintext cocok dengan hash bcrypt.
// Mengembalikan true jika cocok, false jika tidak, dan error jika terjadi masalah saat validasi.
//
// Contoh penggunaan:
//
//	ok, err := CheckPasswordHash("passwordku", hash)
//	fmt.Println(ok, err)
func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
