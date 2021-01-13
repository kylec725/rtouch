# rtouch
rtouch is a recursive touch utility designed to touch existing files, or create
new files while making the required directories for the new file.

rtouch is inspired by [Advanced New File](https://github.com/tanrax/terminal-AdvancedNewFile)

## Installation
Install go for your given distribution.

For Arch Linux users:
```
sudo pacman -S go
```

Set your GOPATH and add your GOPATH bin directory to your system path.

Clone this directory:
```
git clone https://github.com/kylec725/rtouch.git
```

CD into the rtouch directory and run:
```
go install
```


## Usage
```
rtouch path1/file1 path2/file2 ... pathn/filen
```

## Todo
- Add existing touch command flags
