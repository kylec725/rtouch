# rtouch
rtouch is a recursive touch utility designed to touch existing files, or create
new files while making the required directories for the new file.

rtouch is inspired by [Advanced New File](https://github.com/tanrax/terminal-AdvancedNewFile)

## Dependencies
rtouch requires `go`

Install `go` for your given distribution.

For Arch Linux users:
```
sudo pacman -S go
```

## Installation
Clone this directory:
```
git clone https://github.com/kylec725/rtouch.git
```

1. 
- Set your `GOPATH` environment variable and add `$GOPATH/bin` to your system path.

- Go into the `rtouch` directory and run:
```
go install
```

2. 
- Alternatively run:
```
go build
```
- Then place the `rtouch` executable into a directory already within your system path.

## Usage
```
rtouch [path]
```
Terminating a path with / will create/touch a directory instead of a file.

### Examples
Touch a single file
```
rtouch file
```

Touch multiple files
```
rtouch file1 file2
```

Create a new file along with its parent directories
```
rtouch dir1/dir2/dir3/file
```

## Todo
- Add existing touch command flags
