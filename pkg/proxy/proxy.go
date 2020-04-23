package proxy

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/golang/glog"
)

var (
	errProxyTargetFailed = errors.New("[error] proxy target error")
)

// BuildProxy -
func BuildProxy(route *Route) http.Handler {
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			u := req.URL
			glog.V(2).Info("[Proxy Director] url:", u)

			if parsedURL, err := url.ParseRequestURI(req.RequestURI); err == nil {
				u = parsedURL
			}

			if req.URL.Scheme == "" {
				req.URL.Scheme = route.Scheme
			}

			req.URL.Host = route.Host
			req.Host = route.Host

			req.URL.RawPath = u.RawPath
			req.URL.RawQuery = u.RawQuery
			req.RequestURI = ""
		},
		BufferPool: newBufferPool(),
		ModifyResponse: func(resp *http.Response) error {
			glog.V(2).Info("[Proxy ModifyResponse] url", resp.Request.URL.Path, " status code ", resp.StatusCode)

			switch resp.StatusCode {
			case http.StatusSwitchingProtocols, http.StatusOK:
				return nil
			}
			return errProxyTargetFailed
		},
		ErrorHandler: func(w http.ResponseWriter, req *http.Request, err error) {
			glog.Error("[Proxy ErrorHandler] url", req.URL.Path, " --> ", err)

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
