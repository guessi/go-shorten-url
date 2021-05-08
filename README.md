# go-shorten-url

[![Docker Stars](https://img.shields.io/docker/stars/guessi/go-shorten-url.svg)](https://hub.docker.com/r/guessi/go-shorten-url/)
[![Docker Pulls](https://img.shields.io/docker/pulls/guessi/go-shorten-url.svg)](https://hub.docker.com/r/guessi/go-shorten-url/)
[![Docker Automated](https://img.shields.io/docker/automated/guessi/go-shorten-url.svg)](https://hub.docker.com/r/guessi/go-shorten-url/)

a simple shorten url redirection solution with golang


# Prerequisites

- Docker-CE 19.03+
- Docker Compose 1.24.0+


# Usage

    $ docker-compose pull # make sure your image is up-to-date

    $ docker-compose up [-d]

    $ curl http://127.0.0.1:8080/example


# For Developers

    $ go get -u github.com/guessi/go-shorten-url

    $ cd ${GOPATH}/src/github.com/guessi/go-shorten-url

    $ go run main.go --help

    $ go run main.go [-p port] [-c config] [--with-fallback-url] [--no-color] [--debug]


# FAQ

How do I add/remove keywords for url redirection?

    $ vim config/redirections.json

What kind of mobile devices detection are currently supported?

    currently, it only support iOS, AndroidOS

What if there is no "default" action defined?

    it will return "__fallback_url" defined in config/redirections.json

What if there is no device specific action defined?

    it will return url defined in "default" section in config/redirections.json
    or return "fallback_url" defined in config/redirections.js if no device specific url defined

Why is the redirection rules are static?

    it's originally aim to create a simple app with url redirection,
    and it is true that it could be better to integrate with database.
    Pull-Requests are always welcome :-)


# Reference

- [Docker CE](https://www.docker.com/community-edition)
- [Docker Compose](https://docs.docker.com/compose/overview/)
- [Dockerfile Reference](https://docs.docker.com/engine/reference/builder/)


# License

[Apache-2.0](LICENSE)
