GO := GOPATH=$(GOPATH):$(PWD) go

govimda:
	@$(GO) build

clean:
	@$(GO) clean

testpkg:
	@find src -mindepth 1 -type d | while read dir; do \
	  ( cd $$dir; $(GO) test; ); \
	done

test: testpkg
	@$(GO) test

.PHONY: clean testpkg test
