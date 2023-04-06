package trurl_go_test

import (
	trurlgo "github.com/dozheiny/trurl-go"
	"testing"
)

const (
	url = "https://%3a%3amoo%3a%3a@fake.host/hello#frag"
)

func TestParsePort(t *testing.T) {
	t.Parallel()

	res, err := trurlgo.Parse(url)
	if err != nil {
		t.Errorf("Got error on parsing url: %s", err.Error())
	}

	if res.Port != "443" {
		t.Errorf("Exception! port should be %v not %v", 443, res.Port)
	}
}
