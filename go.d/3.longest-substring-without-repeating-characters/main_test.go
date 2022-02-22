package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
    res := lengthOfLongestSubstring("abcabcbb")
    if res != 3  {
        t.Error("Expected 3 ", res)
    }
    res = lengthOfLongestSubstring("bbbbb")
    if res != 1  {
        t.Error("Expected 1 ", res)
    }
    res = lengthOfLongestSubstring("pwwkew")
    if res != 3  {
        t.Error("Expected 3 ", res)
    }
}
