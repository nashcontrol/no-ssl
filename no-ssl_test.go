package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestNoTLSv1_0_TEST1(t *testing.T) {
	in := strings.NewReader("https://www.yahoo.com")
	var out bytes.Buffer
	run(in, &out)

	expectedOutput := "Server https://www.yahoo.com:443 supports TLS 1.0"

	if strings.TrimRight(out.String(), "\n\n") != expectedOutput {
		t.Errorf("expected %s to be equal to %s", out.String(), expectedOutput)
	}
}

func TestNoTLSv1_0_TEST2(t *testing.T) {
	in := strings.NewReader("https://www.facebook.com")
	var out bytes.Buffer
	run(in, &out)

	expectedOutput := "Server https://www.facebook.com:443 supports TLS 1.0"

	if strings.TrimRight(out.String(), "\n\n") != expectedOutput {
		t.Errorf("expected %s to be equal to %s", out.String(), expectedOutput)
	}
}

func TestNoTLSv1_1_TEST1(t *testing.T) {
	in := strings.NewReader("https://stratosphere.lat:443")
	var out bytes.Buffer
	run(in, &out)

	expectedOutput := "Server https://stratosphere.lat:443 supports TLS 1.1"

	if strings.TrimRight(out.String(), "\n\n") != expectedOutput {
		t.Errorf("expected %s to be equal to %s", out.String(), expectedOutput)
	}
}

func TestNoTLSv1_1_TEST2(t *testing.T) {
	in := strings.NewReader("https://www.registrocivil.cl")
	var out bytes.Buffer
	run(in, &out)

	expectedOutput := "Server https://www.registrocivil.cl:443 supports TLS 1.1"

	if strings.TrimRight(out.String(), "\n\n") != expectedOutput {
		t.Errorf("expected %s to be equal to %s", out.String(), expectedOutput)
	}
}

func TestTLSv1_2_TEST1(t *testing.T) {
	in := strings.NewReader("https://www.ssllabs.com")
	var out bytes.Buffer
	run(in, &out)

	expectedOutput := ""

	if strings.TrimRight(out.String(), "\n\n") != expectedOutput {
		t.Errorf("expected %s to support TLS 1.2 and tool reply equal to %s", out.String(), expectedOutput)
	}
}
