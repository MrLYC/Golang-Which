GOENV := env GOPATH=$(GOPATH):`pwd`

.PHONY: build-pkg compile format


compile: format build-pkg
	$(GOENV) go build -o bin/which which
	
format:
	gofmt -w src

build-pkg:
	$(GOENV) go install command
