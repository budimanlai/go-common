package gocommon

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
)

// startTestServer starts a fasthttp server for testing and returns the server, its address, and a close function.
func startTestServer(handler fasthttp.RequestHandler) (addr string, closeFunc func(), err error) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", nil, err
	}
	server := &fasthttp.Server{
		Handler: handler,
	}
	go server.Serve(ln)
	return "http://" + ln.Addr().String(), func() { ln.Close() }, nil
}

func TestHTTPRequestJSON_GET_Success(t *testing.T) {
	handler := func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString(`{"result":"success"}`)
	}
	addr, closeServer, err := startTestServer(handler)
	require.NoError(t, err)
	defer closeServer()

	resp, status, err := HTTPRequest("GET", addr, map[string]string{"Accept": "application/json"}, nil)
	require.NoError(t, err)
	require.Equal(t, fasthttp.StatusOK, status)
	require.JSONEq(t, `{"result":"success"}`, string(resp))
}

func TestHTTPRequestJSON_POST_WithBody(t *testing.T) {
	handler := func(ctx *fasthttp.RequestCtx) {
		body := ctx.PostBody()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		ctx.SetBody(body)
	}
	addr, closeServer, err := startTestServer(handler)
	require.NoError(t, err)
	defer closeServer()

	body := []byte(`{"foo":"bar"}`)
	resp, status, err := HTTPRequest("POST", addr, map[string]string{"Content-Type": "application/json"}, body)
	require.NoError(t, err)
	require.Equal(t, fasthttp.StatusCreated, status)
	require.Equal(t, body, resp)
}

func TestHTTPRequestJSON_Headers(t *testing.T) {
	handler := func(ctx *fasthttp.RequestCtx) {
		val := string(ctx.Request.Header.Peek("X-Custom"))
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString(`{"message":"` + val + `"}`)
	}
	addr, closeServer, err := startTestServer(handler)
	require.NoError(t, err)
	defer closeServer()

	headers := map[string]string{"X-Custom": "header-value"}
	resp, status, err := HTTPRequest("GET", addr, headers, nil)
	require.NoError(t, err)
	require.Equal(t, fasthttp.StatusOK, status)
	require.Equal(t, []byte(`{"message":"header-value"}`), resp)
}

func TestHTTPRequestJSON_Timeout(t *testing.T) {
	handler := func(ctx *fasthttp.RequestCtx) {
		time.Sleep(200 * time.Millisecond)
		ctx.SetStatusCode(fasthttp.StatusOK)
	}
	addr, closeServer, err := startTestServer(handler)
	require.NoError(t, err)
	defer closeServer()

	// Set timeout to 50ms, should timeout
	_, _, err = HTTPRequest("GET", addr, nil, nil, 50)
	require.Error(t, err)
}

func TestHTTPRequestJSON_InvalidURL(t *testing.T) {
	_, _, err := HTTPRequest("GET", "http://invalid.invalid", nil, nil, 100)
	require.Error(t, err)
}
func TestHTTPGetJSON_Success(t *testing.T) {
	handler := func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString(`{"message":"ok"}`)
	}
	addr, closeServer, err := startTestServer(handler)
	require.NoError(t, err)
	defer closeServer()

	resp, status, err := HTTPGetJSON(addr, nil)
	require.NoError(t, err)
	require.Equal(t, fasthttp.StatusOK, status)
	require.JSONEq(t, `{"message":"ok"}`, string(resp))
}

func TestHTTPGetJSON_CustomHeaders(t *testing.T) {
	handler := func(ctx *fasthttp.RequestCtx) {
		accept := string(ctx.Request.Header.Peek("Accept"))
		custom := string(ctx.Request.Header.Peek("X-Test"))
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString(`{"accept":"` + accept + `","custom":"` + custom + `"}`)
	}
	addr, closeServer, err := startTestServer(handler)
	require.NoError(t, err)
	defer closeServer()

	headers := map[string]string{"X-Test": "abc"}
	resp, status, err := HTTPGetJSON(addr, headers)
	require.NoError(t, err)
	require.Equal(t, fasthttp.StatusOK, status)
	require.JSONEq(t, `{"accept":"application/json","custom":"abc"}`, string(resp))
}

func TestHTTPGetJSON_Timeout(t *testing.T) {
	handler := func(ctx *fasthttp.RequestCtx) {
		time.Sleep(200 * time.Millisecond)
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString(`{"message":"delayed"}`)
	}
	addr, closeServer, err := startTestServer(handler)
	require.NoError(t, err)
	defer closeServer()

	// HTTPGetJSON does not accept timeout, but underlying HTTPRequestJSON uses default 10s, so this should succeed
	resp, status, err := HTTPGetJSON(addr, nil)
	require.NoError(t, err)
	require.Equal(t, fasthttp.StatusOK, status)
	require.JSONEq(t, `{"message":"delayed"}`, string(resp))
}

func TestHTTPGetJSON_InvalidURL(t *testing.T) {
	_, _, err := HTTPGetJSON("http://invalid.invalid", nil)
	require.Error(t, err)
}
func TestHTTPRequest_GET_Success(t *testing.T) {
	handler := func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString(`{"status":"ok"}`)
	}
	addr, closeServer, err := startTestServer(handler)
	require.NoError(t, err)
	defer closeServer()

	resp, status, err := HTTPRequest("GET", addr, map[string]string{"Accept": "application/json"}, nil)
	require.NoError(t, err)
	require.Equal(t, fasthttp.StatusOK, status)
	require.JSONEq(t, `{"status":"ok"}`, string(resp))
}

func TestHTTPRequest_POST_WithBody(t *testing.T) {
	handler := func(ctx *fasthttp.RequestCtx) {
		body := ctx.PostBody()
		ctx.SetStatusCode(fasthttp.StatusCreated)
		ctx.SetBody(body)
	}
	addr, closeServer, err := startTestServer(handler)
	require.NoError(t, err)
	defer closeServer()

	body := []byte(`{"foo":"bar"}`)
	resp, status, err := HTTPRequest("POST", addr, map[string]string{"Content-Type": "application/json"}, body)
	require.NoError(t, err)
	require.Equal(t, fasthttp.StatusCreated, status)
	require.Equal(t, body, resp)
}

func TestHTTPRequest_CustomHeaders(t *testing.T) {
	handler := func(ctx *fasthttp.RequestCtx) {
		val := string(ctx.Request.Header.Peek("X-Test"))
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString(`{"header":"` + val + `"}`)
	}
	addr, closeServer, err := startTestServer(handler)
	require.NoError(t, err)
	defer closeServer()

	headers := map[string]string{"X-Test": "my-value"}
	resp, status, err := HTTPRequest("GET", addr, headers, nil)
	require.NoError(t, err)
	require.Equal(t, fasthttp.StatusOK, status)
	require.JSONEq(t, `{"header":"my-value"}`, string(resp))
}

func TestHTTPRequest_Timeout(t *testing.T) {
	handler := func(ctx *fasthttp.RequestCtx) {
		time.Sleep(200 * time.Millisecond)
		ctx.SetStatusCode(fasthttp.StatusOK)
	}
	addr, closeServer, err := startTestServer(handler)
	require.NoError(t, err)
	defer closeServer()

	_, _, err = HTTPRequest("GET", addr, nil, nil, 50)
	require.Error(t, err)
}

func TestHTTPRequest_InvalidURL(t *testing.T) {
	_, _, err := HTTPRequest("GET", "http://invalid.invalid", nil, nil, 100)
	require.Error(t, err)
}

func TestHTTPRequest_DefaultTimeout(t *testing.T) {
	handler := func(ctx *fasthttp.RequestCtx) {
		time.Sleep(50 * time.Millisecond)
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString(`{"msg":"ok"}`)
	}
	addr, closeServer, err := startTestServer(handler)
	require.NoError(t, err)
	defer closeServer()

	// Should succeed because default timeout is 10s
	resp, status, err := HTTPRequest("GET", addr, nil, nil)
	require.NoError(t, err)
	require.Equal(t, fasthttp.StatusOK, status)
	require.JSONEq(t, `{"msg":"ok"}`, string(resp))
}
