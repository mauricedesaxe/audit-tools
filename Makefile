.PHONY: build install

build:
	go build -o audit-tools

install:
	sudo mv audit-tools /usr/local/bin/audit-tools

dev:
	go build -o audit-tools
	sudo mv audit-tools /usr/local/bin/audit-tools