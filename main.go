package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"

    flag "github.com/spf13/pflag"
)

var (
    version = flag.BoolP("version", "v", false, "Show version number")
)

func main() {
    flag.Parse()

    if *version {
        fmt.Println("rtouch 1.00")
        os.Exit(0)
    }

    for _, arg := range os.Args[1:] {
        err := touch(arg)
        if err != nil {
            fmt.Fprintln(os.Stderr, "gotouch:", err)
        }
    }
}

func touch(path string) error {
    var err error

    // Create directories recursively
    if dir := filepath.Dir(path); dir != "" {
        err = os.MkdirAll(dir, 0755)
        if err != nil {
            return err
        }
    }

    // Touch the file or directory
    if _, err = os.Stat(path); os.IsNotExist(err) {
        file, err := os.Create(path)
        if err != nil {
            return err
        }
        file.Close()
    } else {
        now := time.Now().Local()
        err = os.Chtimes(path, now, now)
        if err != nil {
            return err
        }
    }

    return nil
}
