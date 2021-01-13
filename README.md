# rtouch
rtouch is a recursive touch utility designed to touch existing files, or create
new files while making the required directories for the new file.

rtouch is inspired by [Advanced New File](https://github.com/tanrax/terminal-AdvancedNewFile)

## Installation
Install `go` for your given distribution.
For Arch Linux users:
`sudo pacman -S go`

Set your `GOPATH` environment variable and add `$GOPATH/bin` to your system path.

Clone this directory:
`git clone https://github.com/kylec725/rtouch.git`

Go into the `rtouch` directory and run:
`go install`

## Usage
`rtouch [path]`
Terminating a path with / will create/touch a directory instead of a file.

### Examples
Touch a single file
`rtouch file`

Touch multiple files
`rtouch file1 file2`

Create a new file along with its parent directories
`rtouch dir1/dir2/dir3/file`

## Todo
- Add existing touch command flags
