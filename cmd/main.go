package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

type Reader struct {
}

func (r Reader) Read(p []byte) (int, error) {
	for i := 0; i < len(p); i++ {
		p[i] = 0
	}
	return len(p), nil
}
func main() {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var infiniteReader io.Reader = Reader{}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		io.Copy(w, infiniteReader)
	}))
	defer s.Close()
	r, err := http.Get(s.URL)
	if err != nil {
		panic(err)
	}
	if r.StatusCode != http.StatusOK {
		panic(r.Status)
	}
	// b, err := io.ReadAll(r.Body)
	b, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
