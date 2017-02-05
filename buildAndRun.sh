#!/usr/bin/env bash

set -x

# Compile plugins

(cd plugin; docker run --rm -v "$PWD":/usr/src/plugin -w /usr/src/plugin golang:1.8 go build -buildmode=plugin)
(cd plugin2; docker run --rm -v "$PWD":/usr/src/plugin2 -w /usr/src/plugin2 golang:1.8 go build -buildmode=plugin)

# Compile main app

docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.8 go build

# Run it

./myapp plugin/plugin.so plugin2/plugin2.so
