package main

import (
    "fmt"
    "testing"
)

func TestRomaToIntSimple(t *testing.T) {
    result := romeToInt("MCMXCIV")

    fmt.Printf("result %d\n", result)
    //if result != 2 {
    //    t.Error("Expected 1 ", result)
    //}
}