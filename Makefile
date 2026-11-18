VERSION ?= $(shell git describe --tags --always 2>/dev/null | sed 's/^v//' || echo dev)
PREFIX  ?= /usr/local

deb: build
	VERSION=$(VERSION) nfpm package --packager dev

build:
	go build -ldflags "-X main.version=$(VERSION)" -o see .

install: build
	install -D -m 755 see $(DESTDIR)$(PREFIX)/bin/see

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/see