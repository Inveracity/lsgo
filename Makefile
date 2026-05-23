.PHONY: build install

build:
	go build -o lsgo main.go

install: build
	install ./lsgo ~/.local/bin/lsgo
	@rm -f lsgo
	@scripts/alias.sh
