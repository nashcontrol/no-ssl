package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
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

		// normalize input: remove https://, add port 443 if not present, as we are creating a Dialer (TLS connection), not Client object (TLS client side connection)
		if strings.HasPrefix(url, "https") {
			url = string(url[8:])
		}
		if !strings.Contains(url, ":") {
			url += ":443"
		}

		dialer := &net.Dialer{
			Timeout: time.Duration(3 * time.Second),
		}

		// try TLS 1.0 first
		tlsConfig := no_ssl_config(tls.VersionTLS10)
		conn, err := tls.DialWithDialer(dialer, "tcp", url, tlsConfig)

		if conn != nil {
			conn.Close()
		}

		if err == nil {
			fmt.Fprintf(out, "Server https://%s supports TLS 1.0\n", url)
		} else {
			// now try with TLS 1.1
			tlsConfig = no_ssl_config(tls.VersionTLS11)
			conn, err = tls.DialWithDialer(dialer, "tcp", url, tlsConfig)

			if conn != nil {
				conn.Close()
			}

			if err == nil {
				fmt.Fprintf(out, "Server https://%s supports TLS 1.1\n", url)
			}
		}

	}
}

func no_ssl_config(version uint16) *tls.Config {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		MaxVersion:         version,
	}

	return tlsConfig
}
