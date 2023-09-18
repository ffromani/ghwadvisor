#!/bin/bash
#
# look ma, no root!

VERSION=v0.0.2
IMG=quay.io/fromani/ghwadvisor:$VERSION
if [ -n "$MINIMAL" ]; then
	IMG=quay.io/fromani/ghwadvisor-minimal:$VERSION
fi

exec podman run \
  --volume=/sys:/sys:ro \
  --publish=8080:8080 \
  --detach=true \
  --name=ghwadvisor \
  $IMG
