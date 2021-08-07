package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestNoTLS1_1(t *testing.T) {
	in := strings.NewReader("https://www.yahoo.com")
	var out bytes.Buffer
	run(in, &out)

	expectedOutput := "Server https://www.yahoo.com supports TLS 1.0"

	if strings.TrimRight(out.String(), "\n\n") != expectedOutput {
		t.Errorf("expected %s to be equal to %s", out.String(), expectedOutput)
	}
}

func TestNoTLS1_2(t *testing.T) {
	in := strings.NewReader("https://www.facebook.com")
	var out bytes.Buffer
	run(in, &out)

	expectedOutput := "Server https://www.facebook.com supports TLS 1.0"

	if strings.TrimRight(out.String(), "\n\n") != expectedOutput {
		t.Errorf("expected %s to be equal to %s", out.String(), expectedOutput)
	}
}

func TestNoTLS11_1(t *testing.T) {
	in := strings.NewReader("https://kuwaitairways.com")
	var out bytes.Buffer
	run(in, &out)

	expectedOutput := "Server https://kuwaitairways.com supports TLS 1.1"

	if strings.TrimRight(out.String(), "\n\n") != expectedOutput {
		t.Errorf("expected %s to be equal to %s", out.String(), expectedOutput)
	}
}

func TestNoTLS11_2(t *testing.T) {
	in := strings.NewReader("https://www.registrocivil.cl")
	var out bytes.Buffer
	run(in, &out)

	expectedOutput := "Server https://www.registrocivil.cl supports TLS 1.1"

	if strings.TrimRight(out.String(), "\n\n") != expectedOutput {
		t.Errorf("expected %s to be equal to %s", out.String(), expectedOutput)
	}
}

func TestTLS12_1(t *testing.T) {
	in := strings.NewReader("https://www.ssllabs.com")
	var out bytes.Buffer
	run(in, &out)

	expectedOutput := ""

	if strings.TrimRight(out.String(), "\n\n") != expectedOutput {
		t.Errorf("expected %s to support TLS 1.2 and tool reply equal to %s", out.String(), expectedOutput)
	}
}
