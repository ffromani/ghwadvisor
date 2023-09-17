FROM docker.io/golang:1.21 AS builder

WORKDIR /go/src/github.com/ffromani/ghwadvisor
COPY . .

# Build
RUN make

FROM registry.access.redhat.com/ubi9/ubi-minimal
RUN \
    microdnf install -y hwdata && \
    microdnf clean -y all
COPY --from=builder /go/src/github.com/ffromani/ghwadvisor/_out/ghwadvisor /bin/ghwadvisor
EXPOSE 8080/tcp
ENTRYPOINT ["/bin/ghwadvisor"]
