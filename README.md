[![codecov](https://codecov.io/gh/nashcontrol/no-ssl/branch/main/graph/badge.svg?token=YBNYVV3ZSJ)](https://codecov.io/gh/nashcontrol/no-ssl) [![Go Report Card](https://goreportcard.com/badge/github.com/nashcontrol/no-ssl)](https://goreportcard.com/report/github.com/nashcontrol/no-ssl)

# no-ssl

Take a list of domains and probes for support of legacy TLS v1.0/TLS v1.1 protocols and certificate expiration checks.

## Install

```
▶ go get -u github.com/nashcontrol/no-ssl
```

## Basic Usage

no-ssl accepts line-delimited domains (prefixed with HTTPS) on `stdin`. It is best to first confirm host is alive.

```
▶ cat domains.txt
example.com
example.edu
example.net
▶ cat domains.txt | httprobe -s -p https:443 | no-ssl
https://example.com:443 [Certificate Expired, TLS 1.0]
https://example.net:443 [TLS 1.1]
```

It is also possibe to run no-ssl directly with the list of targets, but results may be slower as some hosts might not resolve

```
▶ cat domains.txt | no-ssl
https://example.com:443 [Certificate Expired, TLS 1.0]
https://example.net:443 [TLS 1.1]
```

## Inspired by

1. [httprobe](https://github.com/tomnomnom/httprobe) - Take a list of domains and probe for working http and https servers. 
2. [sec-helpers](https://github.com/vwt-digital/sec-helpers) - Collection of dynamic security related helpers (DAST).
