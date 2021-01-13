package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
)

func main() {
    for _, arg := range os.Args[1:] {
        err := touch(arg)
        if err != nil {
            fmt.Fprintln(os.Stderr, "gotouch:", err)
        }
    }
}

func touch(path string) error {
    var err error
    dir := filepath.Dir(path)

    // Create directories recursively
    if dir != "" {
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
