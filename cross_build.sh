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

cd $(abs_dirname "$0")

GOOS=darwin GOARCH=amd64 go build
mkdir -p bin/osx
mv sshfavo bin/osx/

GOOS=linux GOARCH=amd64 go build
mkdir -p bin/linux
mv sshfavo bin/linux/

GOOS=windows GOARCH=amd64 go build
mkdir -p bin/windows
mv sshfavo.exe bin/windows/
