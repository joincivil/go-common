![Civil Logo](docs/civil_logo_white.png?raw=true)

---
[Civil](https://joincivil.com/) is a decentralized and censorship resistant ecosystem for online Journalism. Read more in our whitepaper.

This repository contains open-source code for common components used in various Civil projects written in Go. 

[![CircleCI](https://img.shields.io/circleci/project/github/joincivil/go-common.svg)](https://circleci.com/gh/joincivil/go-common/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/joincivil/go-common)](https://goreportcard.com/report/github.com/joincivil/go-common)
[![Gitter chat](https://badges.gitter.im/joincivil/Lobby.png)](https://gitter.im/joincivil/Lobby)
[![Telegram chat](https://img.shields.io/badge/chat-telegram-blue.svg)](https://t.me/join_civil)

For Civil's main open-source tools and packages, check out [http://github.com/joincivil/Civil](http://github.com/joincivil/Civil).

## Contributing

Civil's ecosystem is free and open-source, we're all part of it and you're encouraged to be a part of it with us.  We are looking to evolve this into something the community will find helpful and effortless to use.

If you're itching to dwelve deeper inside, [**help wanted**](https://github.com/joincivil/go-common/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22)
and [**good first issue**](https://github.com/joincivil/go-common/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) labels are good places to get started and learn the architecture.

## Install Requirements

This project is using `make` to run setup, builds, tests, etc and has been tested and running on `go 1.11.1`.

Ensure that your `$GOPATH` and `$GOROOT` are setup properly in your shell configuration and that this repo is cloned into the appropriate place in the `$GOPATH`. i.e. `$GOPATH/src/github.com/joincivil/go-common/`

To setup the necessary requirements:

```
make setup
```

### Dependencies

Relies on `dep` [https://golang.github.io/dep/](https://golang.github.io/dep/) for dependency management, updating the `/vendor/` directory in the project.

When adding and removing imports, make sure to run `dep ensure`.  Any adding or removing will require committing the updates on `Gopkg.lock` and `/vendor/` to the repository.

## Lint

Check all the packages for linting errors using a variety of linters via `gometalinter`.  Check the `Makefile` for the up to date list of linters.

```
make lint
```

## Build


```
make build
```

## Testing

Runs the tests and checks code coverage across the project. Produces a `coverage.txt` file for use later.

```
make test
```

## Code Coverage Tool

Run `make test` and launches the HTML code coverage tool.

```
make cover
```