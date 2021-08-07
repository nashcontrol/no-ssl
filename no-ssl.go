package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(in io.Reader, out io.Writer) {

	// get HTTPS urls via stdin
	sc := bufio.NewScanner(in)
	for sc.Scan() {
		url := sc.Text()

		// try TLS 1.0 first
		client := no_ssl(tls.VersionTLS10)
		_, err := client.Get(url)

		// success if there was no TLS connection error, meaning that server does indeed support connection with TLSv1.0 as max version
		if err == nil {
			fmt.Fprintf(out, "Server %s supports TLS 1.0\n", url)
		} else {

			// now try with TLS 1.1
			client := no_ssl(tls.VersionTLS11)
			_, err := client.Get(url)

			if err == nil {

				fmt.Fprintf(out, "Server %s supports TLS 1.1\n", url)
			}
		}
		no_ssl(tls.VersionTLS11)

	}
}

func no_ssl(version uint16) *http.Client {

	// set request timeout
	timeout := time.Duration(2 * time.Second)

	// do not follow redirects
	re := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	// run test on TLS version
	tr := &http.Transport{
		MaxIdleConns:      30,
		IdleConnTimeout:   time.Second,
		DisableKeepAlives: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,

			MaxVersion: version},
		Proxy: http.ProxyFromEnvironment,
	}

	client := &http.Client{
		Transport:     tr,
		CheckRedirect: re,
		Timeout:       timeout,
	}

	return client
}
