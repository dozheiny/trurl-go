package trurl_go

import (
	"net/url"
	"strings"
)

const (
	httpsPort = "443"
	httpPort  = "80"
	https     = "https"
	http      = "http"
)

type Trurl struct {
	Url      string
	Scheme   string
	User     string
	Host     string
	Port     string
	Path     string
	Fragment string
}

func Parse(s string) (*Trurl, error) {
	// Parse URL.
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	t := Trurl{
		Url:  s,
		Port: u.Port(),
	}

	// Parse port.
	if t.Port == "" {
		if strings.Contains(s, https) {
			t.Port = httpsPort
		} else if strings.Contains(s, http) {
			t.Port = httpPort
		}
	}

	return &t, nil
}
