GOENV := env GOPATH=$(GOPATH):`pwd`

.PHONY: compile format

compile:
	$(GOENV) go build -o bin/which which

format:
	gofmt -w src
