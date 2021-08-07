# no-ssl

Take a list of domains and probe for support of TLS 1.0.

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
cat domains.txt | httprobe -s -p https:443 --prefer-https | no-ssl
Server https://example.com supports TLS 1.0
Server https://example.net supports TLS 1.0
```

## Inspired by

1. [httprobe](https://github.com/tomnomnom/httprobe) - Take a list of domains and probe for working http and https servers. 
2. [sec-helpers](https://github.com/vwt-digital/sec-helpers) - Collection of dynamic security related helpers (DAST).