VERSION ?= 0.0.2
REPO ?= quay.io/fromani
IMAGE_TAG_BASE ?= $(REPO)/ghwadvisor
# Image URL to use all building/pushing image targets

CONTAINER_ENGINE ?= podman

all: binaries

outdir:
	mkdir -p _out || :

.PHONY: binaries
binaries: outdir ghwadvisor

ghwadvisor:
	LDFLAGS="-s -w "; \
	go build -mod=vendor -o _out/ghwadvisor -ldflags "$$LDFLAGS" cmd/ghwadvisor/main.go

ghwadvisor.wasm:
	GOOS=wasip1 GOARCH=wasm go build -mod=vendor -o _out/ghwadvisor.wasm cmd/ghwadvisor/main.go

ghwadvisor-static:
	LDFLAGS="-s -w "; \
	CGO_ENABLED=0 go build -mod=vendor -o _out/ghwadvisor -ldflags "$$LDFLAGS" cmd/ghwadvisor/main.go

.PHONY: binaries-static
binaries-static: outdir ghwadvisor-static

.PHONY: binaries-wasm
binaries-wasm: outdir ghwadvisor.wasm

.PHONY: test-unit
test-unit:
	go test ./pkg/...

.PHONY: image
image: container-build
	
.PHONY: container-build
container-build:
	$(CONTAINER_ENGINE) build -t $(IMAGE_TAG_BASE):v$(VERSION) -f Dockerfile .
	$(CONTAINER_ENGINE) build -t $(IMAGE_TAG_BASE)-minimal:v$(VERSION) -f Dockerfile.scratch .

.PHONY: container-build-wasm
container-build-wasm:
	$(CONTAINER_ENGINE) build --annotation run.oci.handler=wasm -t $(IMAGE_TAG_BASE)-wasm:v$(VERSION) -f Dockerfile.wasm .

.PHONY: clean
clean:
	rm -rf _out

.PHONY: gofmt
gofmt:
	@echo "Running gofmt"
	gofmt -s -w `find . -path ./vendor -prune -o -type f -name '*.go' -print`
