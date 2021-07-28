package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestNoTLS1(t *testing.T) {
	in := strings.NewReader("https://www.yahoo.com")
	var out bytes.Buffer
	run(in, &out)

	expectedOutput := "Server https://www.yahoo.com supports TLS 1.0"

	if strings.TrimRight(out.String(), "\n\n") != expectedOutput {
		t.Errorf("expected %s to be equal to %s", out.String(), expectedOutput)
	}
}
