# ws-gui
![go-version](https://img.shields.io/github/go-mod/go-version/cyan903/ws-gui) ![report-card](https://goreportcard.com/badge/github.com/cyan903/ws-gui) ![last-commit](https://img.shields.io/github/last-commit/cyan903/ws-gui)

A simple cross-platform gui for sending/receiving [websockets](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API) using [fyne](https://fyne.io/). Originally made this so I could stop relying on web-based apps as they have issues when making requests to foreign hosts.

![ws-gui-request](https://raw.githubusercontent.com/Cyan903/Static-github/main/ws-gui/ws-gui-request.png)


### To build:
Using [taskfile](https://taskfile.dev/):
```sh
$ task build # use build-dev for unminified version.
$ ./bin/ws.gui
```

### To run:
Make sure you're using go version `1.17` or later.
```
$ go get -u
$ go mod tidy
$ task run
```