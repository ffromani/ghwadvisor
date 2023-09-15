all: binaries

outdir:
	mkdir -p _out || :

.PHONY: binaries
binaries: outdir ghwadvisor

ghwadvisor:
	LDFLAGS="-s -w "; \
	go build -mod=vendor -o _out/ghwadvisor -ldflags "$$LDFLAGS" cmd/ghwadvisor/main.go

.PHONY: test-unit
test-unit:
	go test ./pkg/...

.PHONY: clean
clean:
	rm -rf _out

.PHONY: gofmt
gofmt:
	@echo "Running gofmt"
	gofmt -s -w `find . -path ./vendor -prune -o -type f -name '*.go' -print`
