package trurl_go_test

import (
	trurlgo "github.com/dozheiny/trurl-go"
	"testing"
)

const (
	url = "https://%3a%3amoo%3a%3a@fake.host/hello#frag"
)

func TestParse(t *testing.T) {
	t.Parallel()

	res, err := trurlgo.Parse(url)
	if err != nil {
		t.Errorf("Got error on parsing url: %s", err.Error())
	}

	if res.Port != "443" {
		t.Errorf("Exception! port should be %s not %s", "443", res.Port)
	}

	if res.Host != "fake.host" {
		t.Errorf("Exception! host should be %s not %s", "fake.host", res.Host)
	}

	if res.Url != url {
		t.Errorf("Exception! url should be %s not %s", url, res.Url)
	}

	if res.Scheme != "https" {
		t.Errorf("Exception! scheme should be %s not %s", "https", res.Scheme)
	}

	if res.User != "::moo::" {
		t.Errorf("Exception! user should be %s not %s", "::moo::", res.User)
	}

	if res.Path != "/hello" {
		t.Errorf("Exception! path should be %s not %s", "/hello", res.Path)
	}

	if res.Fragment != "frag" {
		t.Errorf("Exception! frag should be %s not %s", "frag", res.Fragment)
	}

}
