#!/bin/bash
abs_dirname() {
    local cwd="$(pwd)"
    local path="$1"

    while [ -n "$path" ]; do
        cd "${path%/*}"
        local name="${path##*/}"
        path="$(readlink "$name" || true)"
    done

    pwd -P
    cd "$cwd"
}

build_osx() {
    GOOS=darwin GOARCH=amd64 go build
    mkdir -p bin/osx
    mv sshfavo bin/osx/
}

build_linux() {
    GOOS=linux GOARCH=amd64 go build
    mkdir -p bin/linux
    mv sshfavo bin/linux/
}

build_windows() {
    GOOS=windows GOARCH=amd64 go build
    mkdir -p bin/windows
    mv sshfavo.exe bin/windows/
}


cd $(abs_dirname "$0")
build_osx
build_linux
build_windows
