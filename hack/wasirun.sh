#!/bin/bash
# wasirun v0.8.0
exec wasirun --dir /sys --dir /proc  --sockets wasmedgev2 _out/ghwadvisor.wasm
