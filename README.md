## Languages

- English (this README ./)
- Türkçe (coming soon)
- ไทย (coming soon)
- [简体中文](doc/README.zh-CN.md)
- Русский (coming soon)
- Español (coming soon)

# TRON-Vanity-Address-Generator

A lightweight, **open-source** TRON vanity address generator. **Secure**, **offline-only** key generation for custom prefix/suffix TRON addresses. Simple, fast, and developer-friendly.


## Features

- **Fully offline**: No network requests, no telemetry, no auto-upload.
- **Open source**: Small, focused codebase that is easy to audit.
- **CPU-only**: Portable across Linux, macOS, and Windows as a single static binary.
- **Multi-threaded**: Uses all available CPU cores for faster generation.
- **Secure randomness**: Uses cryptographically secure random number generation.

## Download

You can download the latest prebuilt binaries from the GitHub Releases page:

- [Latest release](https://github.com/bigbigha/TRON-Vanity-Address-Generator/releases/latest)
- [TRON Vanity Address Generator v0.1.0](https://github.com/bigbigha/TRON-Vanity-Address-Generator/releases/tag/v0.1.0)


## Helper Scripts (For Non-Developers)

If you prefer not to use the command line directly, you can use the helper scripts:

- [Windows and macOS/Linux helper scripts](scripts/README.md)

## Installation

### For Developers

```bash
# Clone the repository
git clone https://github.com/bigbigha/TRON-Vanity-Address-Generator.git
cd TRON-Vanity-Address-Generator

# Build the binary
go build -o tronvanity ./cmd/tronvanity

# Run the generator
./tronvanity --prefix T888
```

### For Non-Developers (Quick Start)

No need to install Go or Git! Just download the precompiled binary for your system:

1. **Download the binary** from the [GitHub Releases](https://github.com/bigbigha/TRON-Vanity-Address-Generator/releases) page:
   - **Windows**: `tronvanity-windows-amd64.exe`
   - **macOS (Intel)**: `tronvanity-darwin-amd64`
   - **macOS (Apple Silicon)**: `tronvanity-darwin-arm64`
   - **Linux**: `tronvanity-linux-amd64`

2. **Open Terminal or Command Prompt** and navigate to the folder where you downloaded the binary.

3. **Run the tool** with your desired prefix/suffix:

#### Windows:
```cmd
# Example: Generate address with prefix T888
tronvanity-windows-amd64.exe --prefix T888

# Example: Generate address with suffix 9999
tronvanity-windows-amd64.exe --suffix 9999
```

#### macOS:
```bash
# Make the binary executable (only needed once)
chmod +x tronvanity-darwin-amd64  # For Intel Mac
# or
chmod +x tronvanity-darwin-arm64  # For Apple Silicon Mac

# Example: Generate address with prefix T888
./tronvanity-darwin-amd64 --prefix T888

# Example: Generate address with suffix 9999
./tronvanity-darwin-arm64 --suffix 9999
```

#### Linux:
```bash
# Make the binary executable (only needed once)
chmod +x tronvanity-linux-amd64

# Example: Generate address with prefix T888
./tronvanity-linux-amd64 --prefix T888

# Example: Generate address with suffix 9999
./tronvanity-linux-amd64 --suffix 9999
```

**Important**: This tool runs completely offline. You can copy the binary to an air-gapped or isolated computer and use it without any network connection.

## Usage

```bash
# Match prefix only
tronvanity --prefix T888

# Match suffix only
tronvanity --suffix 9999

# Match both prefix and suffix, use 8 workers and stop after 3 matches
tronvanity --prefix T666 --suffix 888 --workers 8 --count 3

# Generate 5 addresses with prefix T123
tronvanity --prefix T123 --count 5
```

### Command line options

- `--prefix`: Prefix to match (e.g. `T888`).
- `--suffix`: Suffix to match (e.g. `9999`).
- `--workers`: Number of worker goroutines (default: number of CPU cores).
- `--count`: Number of matches to find before exiting (default: 1).

## Security tips

1. **Generate offline**: For maximum safety, run this tool on an air-gapped or otherwise offline machine when generating high-value addresses.
2. **Verify addresses**: Always double-check that generated addresses match your desired prefix/suffix before funding them.
3. **Protect private keys**: Store private keys in a secure place (hardware wallet, password manager, or offline encrypted backup). Never paste them into untrusted websites or apps.
4. **Start small**: Before sending large amounts, test with small transactions to confirm that you can successfully receive and spend from the generated address.

## How it works

1. Generates random private keys using Go's cryptographically secure random number generator.
2. Derives the corresponding public key using the secp256k1 elliptic curve (the same curve used by TRON and Ethereum).
3. Computes the TRON address by:
   - Taking the Keccak-256 hash of the uncompressed public key (without the leading `0x04` byte).
   - Taking the last 20 bytes of the hash.
   - Prepending the TRON mainnet prefix `0x41`.
   - Encoding the result using Base58Check to produce a standard `T...` address.
4. Checks whether each generated address matches the specified prefix and/or suffix.
5. Returns all matching addresses along with their corresponding private keys via the CLI.

## Performance

The time required to find a vanity address depends on:

- Length of the prefix/suffix (longer patterns take exponentially more time).
- Exact characters chosen (some Base58 characters make the search space smaller or larger).
- Number of CPU cores and workers you use.
- Pure luck.

Very rough rule of thumb:

- 3-character prefix: seconds to minutes.
- 4-character prefix: minutes to hours.
- 5+ character prefix: hours to days or longer, depending on hardware.

## How it works (short overview)

This tool performs a local brute-force search for TRON vanity addresses by repeatedly generating random key pairs and checking their addresses against your prefix/suffix criteria.

1. It generates a 32-byte private key using Go's cryptographically secure random number generator, then derives the corresponding uncompressed public key (65 bytes) using the secp256k1 elliptic curve (the same curve used by TRON and Ethereum).
2. It computes the Keccak-256 hash of the public key without the leading `0x04` byte, takes the last 20 bytes of the hash, and prepends the TRON mainnet prefix `0x41`.
3. It encodes these 21 bytes using Base58Check to produce a standard `T...` TRON address.
4. Multiple worker goroutines run this process in parallel and check whether each generated address matches the requested prefix and/or suffix. When a match is found, the address and its corresponding private key are sent back to the main process and printed to the CLI.

The entire pipeline runs locally on your machine: no network calls are made, and no private keys are uploaded or shared.

## License

MIT

## Disclaimer

This software is provided "as is", without warranty of any kind. Use it at your own risk. Always verify generated addresses and private keys before using them with real funds, and never store more value in a vanity wallet than you are prepared to secure properly.
>>>>>>> a9e968c (Initial commit)
