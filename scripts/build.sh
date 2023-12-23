#!/bin/bash


readonly CURRENT_DIR_PATH=$(dirname "$0")
readonly REPBASH_DIR_PATH=$(cd "${CURRENT_DIR_PATH}"; cd ../; pwd)
cd "${REPBASH_DIR_PATH}"

readonly remove_past_binary=$(\
  find \
    -mindepth 1 \
    -maxdepth 1 \
    -type f \
    | grep "/repbash-"\
    | awk '{
      if(!$0) next
      print "rm \x22"$0"\x22"
    }'
)
echo "${remove_past_binary}"
case "${remove_past_binary}" in
  "") ;;
  *) bash -c "${remove_past_binary}" ;;
esac

readonly VERSION="0.0.1"

binary_name="repbash-${VERSION}-amd64"
go build -o \
  "${binary_name}" \
  cmd/repbash/main.go | pv
chmod +x "${binary_name}"

binary_name="repbash-${VERSION}-arm64"
GOOS=linux \
GOARCH=arm64 \
CGO_ENABLED=1 \
CC=aarch64-linux-gnu-gcc \
	go build -o \
	  "${binary_name}" \
	  cmd/repbash/main.go | pv
chmod +x "${binary_name}"
