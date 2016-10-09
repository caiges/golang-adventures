package main

import (
	"fmt"
	"strings"
)

func palindrome(source string) bool {
	middle := len(source) / 2

	for i := 0; i <= middle; i++ {
		if source[i] != source[len(source)-i-1] {
			return false
		}
	}

	return true
}

func palindromes(source string, size int) []string {
	candidates := strings.Fields(source)
	found := make([]string, 0)

	for _, candidate := range candidates {
		if palindrome(candidate) {
			found = append(found, candidate)
		}
	}
	return found
}

func main() {
	foundPalindromes := palindromes("tacos are kayak", 5)
	fmt.Println("Found", len(foundPalindromes), "palindromes:")
	for _, v := range foundPalindromes {
		fmt.Println(v)
	}
}
