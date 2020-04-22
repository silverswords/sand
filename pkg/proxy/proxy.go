package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// BuildProxy -
func BuildProxy() *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			u := req.URL

			if parsedURL, err := url.ParseRequestURI(req.RequestURI); err == nil {
				u = parsedURL
			}

			if req.URL.Scheme == "" {
				req.URL.Scheme = "http"
			}

			req.URL.Host = "127.0.0.1:10000"
			req.URL.Path = "/ping"
			req.Host = "127.0.0.1:10000"

			req.URL.RawPath = u.RawPath
			req.URL.RawQuery = u.RawQuery
			req.RequestURI = ""
		},
		BufferPool: newBufferPool(),
	}
}
