#!/bin/bash

# Build script for TRON Vanity Address Generator
# This script builds binaries for all supported platforms

set -e

command -v go >/dev/null 2>&1 || { echo "go not found. Please install Go and ensure it is in PATH."; exit 1; }
command -v git >/dev/null 2>&1 || { echo "git not found. Please install git."; exit 1; }

echo "TRON Vanity Address Generator - Build Script"
echo "============================================"
echo ""

# Create builds directory
mkdir -p builds

# Get version from git tag or use default
VERSION=$(git describe --tags --always 2>/dev/null || echo "v1.0.0")
echo "Building version: $VERSION"
echo ""

# Build for Windows (64-bit)
echo "Building for Windows (amd64)..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o builds/tronvanity-windows-amd64.exe ./cmd/tronvanity

# Build for macOS (Intel)
echo "Building for macOS Intel (amd64)..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o builds/tronvanity-darwin-amd64 ./cmd/tronvanity

# Build for macOS (Apple Silicon)
echo "Building for macOS Apple Silicon (arm64)..."
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o builds/tronvanity-darwin-arm64 ./cmd/tronvanity

# Build for Linux (64-bit)
echo "Building for Linux (amd64)..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o builds/tronvanity-linux-amd64 ./cmd/tronvanity

# Copy helper scripts
echo "Copying helper scripts..."
cp -r scripts/* builds/

# Create SHA256 checksums
echo "Creating checksums..."
cd builds
sha256sum * > SHA256SUMS.txt
cd ..

echo ""
echo "Build completed successfully!"
echo "Binaries are available in the 'builds' directory:"
echo ""
ls -lh builds/
echo ""
echo "SHA256 checksums:"
echo ""
cat builds/SHA256SUMS.txt
echo ""
echo "To create a release, upload the files in the 'builds' directory to GitHub Releases."
