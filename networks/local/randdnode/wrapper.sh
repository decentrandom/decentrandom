#!/usr/bin/env sh

##
## Input parameters
##
BINARY=/randd/${BINARY:-randd}
ID=${ID:-0}
LOG=${LOG:-randd.log}

##
## Assert linux binary
##
if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'randd' E.g.: -e BINARY=randd_my_test_version"
	exit 1
fi
BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"
if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

##
## Run binary with all parameters
##
export RANDDHOME="/randd/node${ID}/randd"

if [ -d "`dirname ${RANDDHOME}/${LOG}`" ]; then
  "$BINARY" --home "$RANDDHOME" "$@" | tee "${RANDDHOME}/${LOG}"
else
  "$BINARY" --home "$RANDDHOME" "$@"
fi

chmod 777 -R /randd