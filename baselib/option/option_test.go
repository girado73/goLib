package option

import "testing"

func TestSome(t *testing.T) {
	got := Some("Hello")
	want := Option[string]{value: "Hello", ok: true}

	if got != want {
		t.Errorf("Some() = %v, want %v", got, want)
	}
}

func TestIsSome(t *testing.T) {
	got := Some("Hello")

	if !got.IsSome() {
		t.Errorf("IsSome() = %v, want true", got.IsSome())
	}
}

func TestIsNone(t *testing.T) {
	got := None[string]()

	if !got.IsNone() {
		t.Errorf("IsNone() = %v, want true", got.IsNone())
	}
}

func TestUnwrap(t *testing.T) {
	got := Some("Hello")
	want := "Hello"

	if got.Unwrap() != want {
		t.Errorf("Unwrap() = %v, want %v", got.Unwrap(), want)
	}
}

func TestUnwrapOr(t *testing.T) {
	got := None[string]()
	want := "Default Value"

	if got.UnwrapOr(want) != want {
		t.Errorf("UnwrapOr() = %v, want %v", got.UnwrapOr(want), want)
	}
}
