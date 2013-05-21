GO := GOPATH=$(GOPATH):$(PWD) go

govimda:
	@$(GO) build

clean:
	@$(GO) clean

.PHONY: clean
