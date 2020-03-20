package server

import (
	"ticker"
	"testing"
	"net/http"
	"net/url"
	"net/http/httptest"
)

func TestServeHTTP_PUT_Happypath(t *testing.T) {
	tk := ticker.NewTicker()
	s := NewServer(tk)
	req := &http.Request{Method: "PUT"}
    req.URL, _ = url.Parse("dummy.com?tickerName=dummy&tickerValue=dummy")
	w := httptest.NewRecorder()

	s.ServeHTTP(w,req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		 t.Errorf("ServeHTTP failed, expected[%d] but got [%d]",http.StatusOK , resp.StatusCode)
	}
}

func TestServeHTTP_PUT_NoTickerNameParameter(t *testing.T) {
	tk := ticker.NewTicker()
	s := NewServer(tk)
	req := &http.Request{Method: "PUT"}
    req.URL, _ = url.Parse("dummy.com?tickerValue=dummy")
	w := httptest.NewRecorder()

	s.ServeHTTP(w,req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		 t.Errorf("ServeHTTP failed, expected[%d] but got [%d]",http.StatusBadRequest , resp.StatusCode)
	}
}

func TestServeHTTP_PUT_NoTickerValueParameter(t *testing.T) {
	tk := ticker.NewTicker()
	s := NewServer(tk)
	req := &http.Request{Method: "PUT"}
    req.URL, _ = url.Parse("dummy.com?tickerName=dummy")
	w := httptest.NewRecorder()

	s.ServeHTTP(w,req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		 t.Errorf("ServeHTTP failed, expected[%d] but got [%d]",http.StatusBadRequest , resp.StatusCode)
	}
}

func TestServeHTTP_DELETE_MethodNotSupported(t *testing.T) {
	tk := ticker.NewTicker()
	s := NewServer(tk)
	req := &http.Request{Method: "DELET"}
    req.URL, _ = url.Parse("dummy.com?tickerName=dummy")
	w := httptest.NewRecorder()

	s.ServeHTTP(w,req)
	resp := w.Result()
	if resp.StatusCode != http.StatusNotFound {
		 t.Errorf("ServeHTTP failed, expected[%d] but got [%d]",http.StatusNotFound , resp.StatusCode)
	}
}
//	body, _ := ioutil.ReadAll(resp.Body)
//	fmt.Println(resp.StatusCode)

