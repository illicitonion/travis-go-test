package lib

import "testing"

func TestFoo(t *testing.T) {
	want := "foo: nil"
	if got := Foo(); got != "foo: nil" {
		t.Fatalf("want %v got %v", want, got)
	}
}
