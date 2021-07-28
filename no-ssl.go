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

	// set request timeout
	timeout := time.Duration(2 * time.Second)

	tr := &http.Transport{
		MaxIdleConns:      30,
		IdleConnTimeout:   time.Second,
		DisableKeepAlives: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,

			MaxVersion: tls.VersionTLS10},
		Proxy: http.ProxyFromEnvironment,
	}

	// do not follow redirects
	re := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	client := &http.Client{
		Transport:     tr,
		CheckRedirect: re,
		Timeout:       timeout,
	}

	// get HTTPS urls via stdin
	sc := bufio.NewScanner(in)
	for sc.Scan() {
		url := sc.Text()

		_, err := client.Get(url)

		if err == nil {
			fmt.Fprintf(out, "Server %s supports TLS 1.0\n", url)
		}

	}
}
