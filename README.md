# filter-kicker

## Description
This filter implements a kicker mechanism to boot out misbehaving clients.


## Features
The filter currently supports:

- non-progressing sessions


## Dependencies
The filter is written in Golang and doesn't have any dependencies beyond the Go extended standard library.

It requires OpenSMTPD 7.5.0 or higher, might work for earlier versions but they are not supported.


## How to install
Install using Go:
```
$ GO111MODULE=on go get github.com/poolpOrg/filter-kicker
$ doas install -m 0555 ~/go/bin/filter-kicker /usr/local/libexec/smtpd/filter-kicker
```

Alternatively, clone the repository, build and install the filter:
```
$ cd filter-kicker/
$ go build
$ doas install -m 0555 filter-kicker /usr/local/libexec/smtpd/filter-kicker
```

On Ubuntu the directory to install to is different:
```
$ sudo install -m 0555 filter-kicker /usr/libexec/opensmtpd/filter-kicker
```


## How to configure
The filter itself requires no configuration.

It must be declared in smtpd.conf and attached to a listener for sessions to go through the kicker:
```
filter "kicker" proc-exec "filter-kicker"

listen on all filter "kicker"
```
