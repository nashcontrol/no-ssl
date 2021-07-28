# no-ssl

Take a list of domains and probe for support of TLS 1.0.

## Install

```
▶ go get -u github.com/nashcontrol/no-ssl
```

## Basic Usage

no-ssl accepts line-delimited domains (prefixed with HTTPS) on `stdin`:

```
▶ cat domains.txt
example.com
example.edu
example.net
cat domains.txt | httprobe --prefer-https | no-ssl
Server https://example.com supports TLS 1.0
Server https://example.net supports TLS 1.0
```