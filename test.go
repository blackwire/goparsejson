package main

import (
    "./json"
    "fmt"
)

func main() {
    data, _ := json.FromFile("data.json")
    fmt.Printf("Float: %T\n", data.Get("array.float"))
    fmt.Printf("Integer: %T\n", data.Get("array.int"))
    fmt.Printf("String: %T\n", data.Get("string"))
    fmt.Printf("Map: %T\n", data.Get("array"))
}