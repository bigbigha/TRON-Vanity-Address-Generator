# Helper Scripts for TRON Vanity Address Generator

These scripts provide a simple interface for non-technical users to generate TRON vanity addresses without using the command line directly.

## Windows Script

### Usage
1. Download `tronvanity-windows-amd64.exe` from the [GitHub Releases](https://github.com/bigbigha/TRON-Vanity-Address-Generator/releases) page
2. Download `windows-run.bat` from this folder
3. Place both files in the same folder
4. Double-click `windows-run.bat` to run the script
5. Follow the on-screen instructions to enter your desired prefix, suffix, and number of addresses

### Features
- Simple text-based interface
- No command line knowledge required
- Automatically repeats until you press Ctrl+C to exit

## macOS/Linux Script

### Usage
1. Download the appropriate binary for your system from the [GitHub Releases](https://github.com/bigbigha/TRON-Vanity-Address-Generator/releases) page:
   - `tronvanity-darwin-amd64` (Intel Mac)
   - `tronvanity-darwin-arm64` (Apple Silicon Mac)
   - `tronvanity-linux-amd64` (Linux)
2. Download `unix-run.sh` from this folder
3. Place both files in the same folder
4. Open Terminal and navigate to the folder
5. Run: `chmod +x unix-run.sh`
6. Run: `./unix-run.sh`
7. Follow the on-screen instructions

### Features
- Automatic detection of your operating system and architecture
- Simple text-based interface
- No advanced command line knowledge required
- Automatically makes the binary executable
- Automatically repeats until you press Ctrl+C to exit

## Security Note

These scripts are simple wrappers that only call the local binary. They do not:
- Access the internet
- Upload any data
- Install any software
- Modify any system files

The scripts and the main binary run completely offline on your local machine.