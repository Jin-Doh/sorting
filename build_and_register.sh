#!/bin/bash

set -e

# The name of the binary to be built
BIN_NAME="sorting"

# Build the main.go in the current directory
go build -o "$BIN_NAME" main.go

# Move the binary to /usr/local/bin (may require sudo)
if [ -w /usr/local/bin ]; then
    mv -f "$BIN_NAME" /usr/local/bin/
else
    sudo mv -f "$BIN_NAME" /usr/local/bin/
fi

echo "Installation complete: /usr/local/bin/$BIN_NAME"
echo "You can run it using the command '$BIN_NAME' in the terminal."
