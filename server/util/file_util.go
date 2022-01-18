package util

import (
    "os"
    "errors"
    "fmt"
)

func IsFileExists(path string) bool {
    exist := false
    if _, err := os.Stat(path); err == nil {
        exist = true
    } else if errors.Is(err, os.ErrNotExist) {
        exist = false
    } else {
        fmt.Printf("os.Stat Error is : %s.\n", err)
        exist = false
    }
    return exist
}
