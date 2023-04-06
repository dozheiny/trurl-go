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
	Password string
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
		Url:      s,
		Port:     u.Port(),
		Host:     u.Host,
		Scheme:   u.Scheme,
		User:     u.User.Username(),
		Path:     u.Path,
		Fragment: u.Fragment,
	}

	// Parse Password.
	if pass, ok := u.User.Password(); ok {
		t.Password = pass
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
