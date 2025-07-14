package gocommon

import (
	"time"
)

// TimeToString mengonversi objek time.Time ke string dengan format "2006-01-02 15:04:05".
// Fungsi ini berguna untuk menyimpan atau menampilkan waktu dalam format yang mudah dibaca.
// Jika objek time.Time tidak valid, fungsi ini akan mengembalikan string kosong.
//
// Contoh penggunaan:
// t := time.Now()
// formattedTime := TimeToString(t)
// fmt.Println(formattedTime) // Output: "2023-10-01 12:34:56" (contoh, tergantung waktu saat ini)
func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// StringToTime mengonversi string dengan format "2006-01-02 15:04:05" ke objek time.Time.
// Jika string tidak sesuai format, fungsi ini akan mengembalikan zero time dan error.
//
// Contoh penggunaan:
// s := "2023-10-01 12:34:56"
// t, err := StringToTime(s)
//
//	if err != nil {
//	    fmt.Println("Format waktu tidak valid:", err)
//	} else {
//
//	    fmt.Println(t)
//	}
func StringToTime(s string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", s)
}

// ConvertToLocalTime mengonversi objek time.Time ke local timezone sesuai setting server.
// Fungsi ini berguna jika waktu yang diterima masih dalam UTC atau timezone lain.
//
// Contoh penggunaan:
// t := time.Now().UTC()
// localTime := ConvertToLocalTime(t)
// fmt.Println(localTime)
func ConvertToLocalTime(t time.Time) time.Time {
	return t.In(time.Local)
}

// StringWithTZToLocalTime mengonversi string datetime dengan timezone ke local timezone sesuai setting server.
// Format string yang didukung: "2006-01-02 15:04:05-07:00" atau "2006-01-02T15:04:05Z07:00".
// Jika parsing gagal, akan mengembalikan zero time dan error.
//
// Contoh penggunaan:
// s := "2023-10-01T12:34:56+07:00"
// localTime, err := StringWithTZToLocalTime(s)
//
//	if err != nil {
//	    fmt.Println("Format waktu tidak valid:", err)
//	} else {
//
//	    fmt.Println(localTime)
//	}
func StringWithTZToLocalTime(s string) (time.Time, error) {
	layouts := []string{
		time.RFC3339,                // "2006-01-02T15:04:05Z07:00"
		"2006-01-02 15:04:05-07:00", // "2006-01-02 15:04:05-07:00"
	}
	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, s)
		if err == nil {
			return t.In(time.Local), nil
		}
	}
	return time.Time{}, err
}

// GetCurrentTimeInLocalZone mengembalikan waktu saat ini dalam local timezone sesuai setting server.
// Fungsi ini berguna untuk mendapatkan waktu sekarang yang sesuai dengan zona waktu server.
//
// Contoh penggunaan:
//
// currentTime := GetCurrentTimeInLocalZone()
//
// fmt.Println(currentTime) // Output: "2023-10-01 12:34:56" (contoh, tergantung waktu saat ini)
func GetCurrentTimeInLocalZone() time.Time {
	return time.Now().In(time.Local)
}
