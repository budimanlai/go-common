package gocommon

import (
	"time"

	"encoding/json"

	"github.com/valyala/fasthttp"
)

// HTTPRequest mengirim permintaan HTTP dengan metode, URL, header, dan body yang ditentukan,
// serta mengharapkan respons dalam format JSON. Fungsi ini menggunakan pustaka fasthttp untuk performa tinggi.
//
// Parameter:
//   - method: Metode HTTP (misal: "GET", "POST").
//   - url: URL tujuan permintaan.
//   - headers: Map berisi pasangan key-value header yang akan disertakan dalam permintaan.
//   - body: Body permintaan dalam bentuk byte slice. Bisa nil untuk metode seperti GET.
//   - timeout: Timeout opsional dalam milidetik. Jika tidak diberikan, default 10.000 ms (10 detik).
//
// Return:
//   - []byte: Body respons dalam bentuk byte slice.
//   - int: Kode status HTTP dari respons.
//   - error: Error jika permintaan gagal atau respons bukan JSON valid.
//
// Fungsi ini memvalidasi bahwa body respons adalah JSON yang valid. Jika tidak, akan mengembalikan error.
func HTTPRequest(method, url string, headers map[string]string, body []byte, timeout ...int) ([]byte, int, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if body != nil {
		req.SetBody(body)
	}

	// Default timeout: 10 seconds
	to := 10 * 1000 // milliseconds
	if len(timeout) > 0 {
		to = timeout[0]
	}

	err := fasthttp.DoTimeout(req, resp, time.Duration(to)*time.Millisecond)
	if err != nil {
		return nil, 0, err
	}

	responseBody := make([]byte, len(resp.Body()))
	copy(responseBody, resp.Body())

	return responseBody, resp.StatusCode(), nil
}

// HTTPGetJSON mengirim permintaan HTTP GET ke URL yang ditentukan dengan header yang diberikan,
// serta menambahkan header "Accept" dengan nilai "application/json". Fungsi ini mengembalikan body respons sebagai byte slice,
// kode status HTTP, dan error jika permintaan gagal.
//
// Contoh penggunaan:
//
//	resp, status, err := HTTPGetJSON("https://api.example.com/data", nil)
//	if err != nil {
//	    // tangani error
//	}
//	fmt.Printf("Status: %d, Body: %s\n", status, string(resp))
func HTTPGetJSON(url string, headers map[string]string) ([]byte, int, error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Accept"] = "application/json"
	resp, status, err := HTTPRequest("GET", url, headers, nil)
	if err != nil {
		return nil, status, err
	}

	return resp, status, err
}

// HTTPPostJSON mengirim permintaan HTTP POST dengan body yang dienkode dalam format JSON ke URL yang ditentukan.
// Fungsi ini mengatur header "Content-Type" dan "Accept" menjadi "application/json".
// Body yang diberikan akan di-marshal ke JSON, permintaan dikirim, dan respons divalidasi agar merupakan JSON yang valid.
// Fungsi mengembalikan body respons sebagai byte slice, kode status HTTP, dan error jika terjadi kesalahan.
//
// Parameter:
//   - url: Endpoint tujuan permintaan POST.
//   - headers: Header HTTP opsional yang akan disertakan dalam permintaan. Jika nil, akan dibuat map baru.
//   - body: Data yang akan dienkode ke JSON dan dikirim sebagai body permintaan.
//
// Return:
//   - []byte: Body respons.
//   - int: Kode status HTTP.
//   - error: Error jika permintaan gagal, body gagal di-marshal, atau respons bukan JSON valid.
func HTTPPostJSON(url string, headers map[string]string, body interface{}, timeout ...int) ([]byte, int, error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, 0, err
	}

	resp, status, err := HTTPRequest("POST", url, headers, jsonBody, timeout...)
	if err != nil {
		return nil, status, err
	}

	return resp, status, nil
}
