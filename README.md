# Simple Go SOCKS5 proxy #

[![Build Status][build-img]][build] [![Godoc][godoc-img]][godoc]

[build]: https://github.com/StalkR/socks-proxy/actions/workflows/build.yml
[build-img]: https://github.com/StalkR/socks-proxy/actions/workflows/build.yml/badge.svg
[godoc]: https://godoc.org/github.com/StalkR/socks-proxy
[godoc-img]: https://godoc.org/github.com/StalkR/socks-proxy?status.png

A simple SOCKS5 proxy in Go, no caching.

It listens on TCP IPv4/IPv6 at the specified port.

Example:

    $ go run socks_proxy.go -listen :1080

# Setup #

Install go package, create Debian package, install:

    $ go get -u github.com/StalkR/socks-proxy
    $ cd $GOPATH/src/github.com/StalkR/socks-proxy
    $ fakeroot debian/rules clean binary
    $ sudo dpkg -i ../socks-proxy_1-1_amd64.deb

Configure in `/etc/default/socks-proxy` and start with `/etc/init.d/socks-proxy start`.

# License #

[Apache License, version 2.0](http://www.apache.org/licenses/LICENSE-2.0).

# Thanks #

- the powerful [go-socks5](https://github.com/armon/go-socks5) library by [Armon Dadgar](https://github.com/armon)

# Bugs, feature requests, questions #

Create a [new issue](https://github.com/StalkR/socks-proxy/issues/new).
