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
	RawUrl   *url.URL
}

// Parse will give s as string and parse it into Trurl struct.
func Parse(s string) *Trurl {
	// Parse URL.
	u, err := url.Parse(s)
	if err != nil {
		return &Trurl{}
	}

	t := Trurl{
		Url:      strings.ToLower(s),
		Port:     u.Port(),
		Host:     u.Host,
		Scheme:   u.Scheme,
		User:     u.User.Username(),
		Path:     u.Path,
		Fragment: u.Fragment,
		RawUrl:   u,
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

	return &t
}

// SetHost will set host on given url and Trurl.Host value.
func (t *Trurl) SetHost(h string) *Trurl {
	t.Host = h
	t.RawUrl.Host = h
	t.Url = strings.ToLower(t.RawUrl.String())

	return t
}
