package string

import (
	array "girado73/goLib/baselib/array"
	"testing"
)

func TestChars(t *testing.T) {
	got := Chars("Hello")
	want := []rune{'H', 'e', 'l', 'l', 'o'}

	if !array.Compare(got, want) {
		t.Errorf("Chars() = %v, want %v", got, want)
	}
}

func TestChunkSplit(t *testing.T) {
	got := ChunkSplit("HelloWorld", 5, "-")
	want := "Hello-World"

	if got != want {
		t.Errorf("ChunkSplit() = %v, want %v", got, want)
	}
}

func TestJoinRunes(t *testing.T) {
	got := JoinRunes([]rune{'H', 'e', 'l', 'l', 'o'})
	want := "Hello"

	if got != want {
		t.Errorf("JoinRunes() = %v, want %v", got, want)
	}
}

func TestTrim(t *testing.T) {
	got := Trim("  Hello World  ")
	want := "HelloWorld"

	if got != want {
		t.Errorf("Trim() = %v, want %v", got, want)
	}
}

func TestStrToUpper(t *testing.T) {
	got := StrToUpper("Hello World")
	want := "HELLO WORLD"

	if got != want {
		t.Errorf("StrToUpper() = %v, want %v", got, want)
	}
}

func TestStrToLower(t *testing.T) {
	got := StrToLower("Hello World")
	want := "hello world"

	if got != want {
		t.Errorf("StrToLower() = %v, want %v", got, want)
	}
}
