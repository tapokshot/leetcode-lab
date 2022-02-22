package main

import "testing"

func TestTwoSumBruteForce(t *testing.T) {
    input := []int{2, 7, 11, 15}
    target := 9
    result := twoSumBruteForce(input, target)
    if result[0] != 0 || result[1] != 1 {
       t.Error("Expected [0,1] ", result)
    }
    result = onePassHashTable(input, target)
    if len(result) == 0 || result[0] != 0 || result[1] != 1 {
        t.Error("Expected [0,1] ", result)
    }
    result = twoPassHashTable(input, target)
    if len(result) == 0 || result[0] != 0 || result[1] != 1 {
        t.Error("Expected [0,1] ", result)
    }


}
