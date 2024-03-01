package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    f, err := os.Create("testFile.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    fmt.Println(f.Name())
}
