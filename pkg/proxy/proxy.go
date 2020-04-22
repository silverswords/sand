package proxy

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	errProxyTargetFailed = errors.New("[error] proxy target error")
)

// BuildProxy -
func BuildProxy(route *Route) http.Handler {
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			u := req.URL

			if parsedURL, err := url.ParseRequestURI(req.RequestURI); err == nil {
				u = parsedURL
			}

			if req.URL.Scheme == "" {
				req.URL.Scheme = "http"
			}

			req.URL.Host = route.Host[0]
			req.Host = route.Host[0]

			req.URL.RawPath = u.RawPath
			req.URL.RawQuery = u.RawQuery
			req.RequestURI = ""
		},
		BufferPool: newBufferPool(),
		ModifyResponse: func(resp *http.Response) error {
			if resp.StatusCode != http.StatusOK {
				return errProxyTargetFailed
			}

			return nil
		},
		ErrorHandler: func(w http.ResponseWriter, req *http.Request, err error) {
			statusCode := http.StatusInternalServerError

			switch {
			case err == errProxyTargetFailed:
				statusCode = http.StatusBadGateway
			default:
				if e, ok := err.(net.Error); ok {
					if e.Timeout() {
						statusCode = http.StatusGatewayTimeout
					} else {
						statusCode = http.StatusBadGateway
					}
				}
			}

			w.WriteHeader(statusCode)
		},
	}
}
