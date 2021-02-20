package middleware

import (
	"net/http"
	"strings"
	"io"
	"compress/gzip"
)

type GzipMiddleware struct {
	Next http.Handler
}

func (gm GzipMiddleware) validateAcceptGzipEnconding(w *http.ResponseWriter, r *http.Request) bool {
	encoding := r.Header.Get("Accept-Encoding")
	if !strings.Contains(encoding, "gzip") {
		gm.Next.ServeHTTP(*w, r)
		return false;
	}

	return true;
}

func (gm GzipMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if gm.Next == nil {
		gm.Next = http.DefaultServeMux
	}

	if acceptsEncoding := gm.validateAcceptGzipEnconding(&w, r); acceptsEncoding == true {
		w.Header().Add("Content-Encoding", "gzip")
		gzipwriter := gzip.NewWriter(w)
		defer gzipwriter.Close()
		grw := gzipReponseWriter {
			ResponseWriter: w,
			Writer: gzipwriter,
		}
		gm.Next.ServeHTTP(grw, r)
	}

}

type gzipReponseWriter struct {
	http.ResponseWriter
	io.Writer
}

func (grw gzipReponseWriter) Write(data []byte) (int, error) {
	return grw.Writer.Write(data);
}