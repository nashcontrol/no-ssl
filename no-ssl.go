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

		// normalize input: remove https://, add port 443 if not present, strip trailing "/" - as we are creating a Dialer (TLS connection), not Client object (TLS client side connection)
		if strings.HasPrefix(url, "https") {
			url = string(url[8:])
		}

		if strings.HasSuffix(url, "/") {
			url = strings.TrimSuffix(url, "/")

		}

		if !strings.Contains(url, ":") {
			url += ":443"
		}

		dialer := &net.Dialer{
			Timeout: time.Duration(3 * time.Second),
		}

		// try TLS 1.0 first
		tlsConfig := configTLS(tls.VersionTLS10)
		conn, err := tls.DialWithDialer(dialer, "tcp", url, tlsConfig)

		// aggragete results
		var assetResult []string

		if conn != nil {
			expireDate := conn.ConnectionState().PeerCertificates[0].NotAfter
			if expireDate.Before(time.Now()) {
				assetResult = append(assetResult, "Certificate Expired")
			}
			conn.Close()
		}

		if err == nil {
			assetResult = append(assetResult, "TLS 1.0")
		} else {
			// now try with TLS 1.1
			tlsConfig = configTLS(tls.VersionTLS11)
			conn, err = tls.DialWithDialer(dialer, "tcp", url, tlsConfig)

			if conn != nil {
				conn.Close()
			}

			if err == nil {
				assetResult = append(assetResult, "TLS 1.1")
			}
		}

		//if we have any gaps:
		if len(assetResult) > 0 {
			fmt.Fprintf(out, "https://%s [%s]\n", url, strings.Join(assetResult[:], ", "))
		}

	}
}

func configTLS(version uint16) *tls.Config {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		MaxVersion:         version,
	}

	return tlsConfig
}
