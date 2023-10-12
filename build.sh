#!/bin/bash

BBC_VER="v0.0.1-beta-1"

DIST_DIR=".dist"

declare -a os_array=("linux" "darwin" "windows")
declare -a arch_array=("amd64" "arm64")

mkdir -p "$DIST_DIR"

for OS in "${os_array[@]}"; do
    for ARCH in "${arch_array[@]}"; do
        BBC_DIST_PATH="$DIST_DIR/bosh-cli-completion-$BBC_VER-$OS-$ARCH"
        echo "build $BBC_DIST_PATH"
        if GOOS=$OS GOARCH=$ARCH go build -o "$BBC_DIST_PATH"  >> "$BBC_DIST_PATH.log"; then
            sha256sum "$BBC_DIST_PATH"  | awk '{print $1}' > "$BBC_DIST_PATH.sum"
        fi
    done
done

# build a standalone, statically-linked binary using an external linker
OS="linux"
ARCH="amd64"
BBC_DIST_PATH="$DIST_DIR/bosh-cli-completion-standalone-$BBC_VER-$OS-$ARCH"
echo "build $BBC_DIST_PATH"
if GOOS=$OS GOARCH=$ARCH go build -ldflags '-linkmode external -extldflags "-static"' -o "$BBC_DIST_PATH" >> "$BBC_DIST_PATH.log"; then
    sha256sum "$BBC_DIST_PATH"  | awk '{print $1}' > "$BBC_DIST_PATH.sum"
fi
