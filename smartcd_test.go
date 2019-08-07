package main

import (
	"fmt"
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func Test01(t *testing.T) {
	result := shorten("abc-def-ghi")

	if !contains(result, "abc-def-ghi") {
		t.Fatal("failed test")
	}
	if !contains(result, "abcdefghi") {
		t.Fatal("failed test")
	}
	if !contains(result, "a-d-g") {
		t.Fatal("failed test")
	}
	if !contains(result, "adg") {
		t.Fatal("failed test")
	}
	if !contains(result, "ab-de-gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdegh") {
		t.Fatal("failed test")
	}
}

func Test02(t *testing.T) {
	result := shorten("ab-de-gh")

	if !contains(result, "ab-de-gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdegh") {
		t.Fatal("failed test")
	}
	if !contains(result, "a-d-g") {
		t.Fatal("failed test")
	}
	if !contains(result, "adg") {
		t.Fatal("failed test")
	}
	if !contains(result, "ab-de-gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdegh") {
		t.Fatal("failed test")
	}
}

func Test03(t *testing.T) {
	result := shorten("abc-d-ghi")

	fmt.Println(result)

	if !contains(result, "ab-d-gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdgh") {
		t.Fatal("failed test")
	}
	if !contains(result, "a-d-g") {
		t.Fatal("failed test")
	}
	if !contains(result, "adg") {
		t.Fatal("failed test")
	}
	if !contains(result, "ab-d-gh") {
		t.Fatal("failed test")
	}
	if !contains(result, "abdgh") {
		t.Fatal("failed test")
	}
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
