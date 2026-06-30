all: build symlink

build:
	go build -o ncah cmd/ncah/main.go

symlink:
	sudo ln -sf $(shell pwd)/ncah /usr/bin/ncah
