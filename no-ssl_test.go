package main

import (
	"bytes"
	"crypto/tls"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestTLSv1_1_Self(t *testing.T) {

	certPem := []byte(`-----BEGIN CERTIFICATE-----
MIIBhTCCASugAwIBAgIQIRi6zePL6mKjOipn+dNuaTAKBggqhkjOPQQDAjASMRAw
DgYDVQQKEwdBY21lIENvMB4XDTE3MTAyMDE5NDMwNloXDTE4MTAyMDE5NDMwNlow
EjEQMA4GA1UEChMHQWNtZSBDbzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABD0d
7VNhbWvZLWPuj/RtHFjvtJBEwOkhbN/BnnE8rnZR8+sbwnc/KhCk3FhnpHZnQz7B
5aETbbIgmuvewdjvSBSjYzBhMA4GA1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggr
BgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1UdEQQiMCCCDmxvY2FsaG9zdDo1
NDUzgg4xMjcuMC4wLjE6NTQ1MzAKBggqhkjOPQQDAgNIADBFAiEA2zpJEPQyz6/l
Wf86aX6PepsntZv2GYlA5UpabfT2EZICICpJ5h/iI+i341gBmLiAFQOyTDT+/wQc
6MF9+Yw1Yy0t
-----END CERTIFICATE-----`)
	keyPem := []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIIrYSSNQFaA2Hwf1duRSxKtLYX5CB04fSeQ6tF1aY/PuoAoGCCqGSM49
AwEHoUQDQgAEPR3tU2Fta9ktY+6P9G0cWO+0kETA6SFs38GecTyudlHz6xvCdz8q
EKTcWGekdmdDPsHloRNtsiCa697B2O9IFA==
-----END EC PRIVATE KEY-----`)

	cert, err := tls.X509KeyPair(certPem, keyPem)

	if err != nil {
		log.Fatal(err)
	}

	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS11,
	}
	srv := &http.Server{
		TLSConfig:    cfg,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		Addr:         "localhost:2000",
	}

	go srv.ListenAndServeTLS("", "")

	defer srv.Close()
	time.Sleep(1500 * time.Millisecond)

	in := strings.NewReader("https://localhost:2000")
	var out bytes.Buffer
	run(in, &out)

	srv.Close()

	expectedOutput := "https://localhost:2000 [TLS 1.1]"

	if strings.TrimRight(out.String(), "\n\n") != expectedOutput {
		t.Errorf("expected %s to be equal to %s", out.String(), expectedOutput)
	}
}
func TestNoTLSv1_0_TEST1(t *testing.T) {
	in := strings.NewReader("https://www.yahoo.com")
	var out bytes.Buffer
	run(in, &out)

	expectedOutput := "https://www.yahoo.com:443 [TLS 1.0]"

	if strings.TrimRight(out.String(), "\n\n") != expectedOutput {
		t.Errorf("expected %s to be equal to %s", out.String(), expectedOutput)
	}
}

func TestNoTLSv1_0_TEST2(t *testing.T) {
	in := strings.NewReader("https://www.facebook.com")
	var out bytes.Buffer
	run(in, &out)

	expectedOutput := "https://www.facebook.com:443 [TLS 1.0]"

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
