#!/bin/bash
#
# look ma, no root!

VERSION=v0.0.1
exec podman run \
  --volume=/sys:/sys:ro \
  --publish=8080:8080 \
  --detach=true \
  --name=ghwadvisor \
  quay.io/fromani/ghwadvisor:$VERSION
