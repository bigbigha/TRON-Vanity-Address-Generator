#!/bin/bash

clear
echo "TRON Vanity Address Generator"
echo "============================"
echo ""

# Detect the binary name based on the system
BINARY_NAME=""
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    if [[ $(uname -m) == "arm64" ]]; then
        BINARY_NAME="tronvanity-darwin-arm64"
    else
        BINARY_NAME="tronvanity-darwin-amd64"
    fi
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    # Linux
    BINARY_NAME="tronvanity-linux-amd64"
else
    echo "Unsupported operating system."
    exit 1
fi

# Check if binary exists
if [ ! -f "./$BINARY_NAME" ]; then
    echo "Error: $BINARY_NAME not found in current directory."
    echo "Please make sure the binary is in the same folder as this script."
    exit 1
fi

# Make binary executable
chmod +x "./$BINARY_NAME"

while true; do
    echo "Please enter the prefix you want to match (or press Enter to skip):"
    read prefix
    echo ""

    echo "Please enter the suffix you want to match (or press Enter to skip):"
    read suffix
    echo ""

    echo "Please enter the number of addresses to generate (default: 1):"
    read count
    if [ -z "$count" ]; then
        count=1
    fi
    echo ""

    echo "Searching for TRON vanity addresses..."
    echo ""

    # Build the command
    cmd="./$BINARY_NAME"
    if [ ! -z "$prefix" ]; then
        cmd="$cmd --prefix $prefix"
    fi
    if [ ! -z "$suffix" ]; then
        cmd="$cmd --suffix $suffix"
    fi
    cmd="$cmd --count $count"

    # Execute the command
    $cmd

    echo ""
    echo ""
    echo "Press Enter to generate more addresses, or Ctrl+C to exit..."
    read -r
done