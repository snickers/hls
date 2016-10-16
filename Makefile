.PHONY: all run build help 

help:
	@echo '    build ...................... go get the dependencies'

build:
	go get

run:
	go run main.go
