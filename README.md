# go-common

Please use https://github.com/budimanlai/go-pkg

Kumpulan utilitas dan helper Go yang dapat digunakan kembali untuk tugas-tugas umum dalam pengembangan web, pemrosesan data, dan lainnya.

## Fitur
- Utilitas tanggal dan waktu
- Helper manipulasi string
- Utilitas HTTP (menggunakan [fasthttp](https://github.com/valyala/fasthttp))
- Helper database (menggunakan [sqlx](https://github.com/jmoiron/sqlx))
- Utilitas validasi (menggunakan [validator](https://github.com/go-playground/validator))

## Instalasi

```
go get github.com/budimanlai/go-common
```

## Contoh Penggunaan

```go
import (
    "github.com/budimanlai/go-common"
)

// Contoh penggunaan helper tanggal
common.FormatDate(time.Now())
```

Lihat direktori `examples/` untuk contoh penggunaan lainnya.

## Struktur Direktori
- `app.go`, `date.go`, `helpers.go`, `http.go`, `strings.go`: File utilitas utama
- `models/`: Model data
- `examples/`: Contoh penggunaan

## Dependensi
- fasthttp
- sqlx
- validator
- Lihat `go.mod` untuk daftar lengkap

## Lisensi
MIT
