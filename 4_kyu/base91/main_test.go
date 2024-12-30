package main

import "testing"

func TestEncode(t *testing.T) {
	first := Encode([]byte("test"))
	if string(first) != "fPNKd" {
		t.Fatalf("Encoded 'test' should be 'fPNKd', isntead got=%v", string(first))
	}
	second := Encode([]byte("Hello World!"))
	if string(second) != ">OwJh>Io0Tv!8PE" {
		t.Fatalf("Encoded 'Hello World!' should be '>OwJh>Io0Tv!8PE', isntead got=%v", string(second))
	}
	third := Encode([]byte("13901195017114136957958115016332136828968799116053"))
	if string(third) != "'1CF2i%zgJR*9NbRkwuxVoyoG3F?lN$SgwZ.Ak&GAJG?SCLUK)gbWo$GwJ::WC" {
		t.Fatalf("Decoded '13901195017114136957958115016332136828968799116053' should be '1CF2i%%zgJR*9NbRkwuxVoyoG3F?lN$SgwZ.Ak&GAJG?SCLUK)gbWo$GwJ::WC instead got=%v", string(third))
	}
}
func TestDecode(t *testing.T) {
	first := Decode([]byte("fPNKd"))
	if string(first) != "test" {
		t.Fatalf("Decoded 'fPNKd' should be 'test', isntead got=%v", string(first))
	}
	second := Decode([]byte(">OwJh>Io0Tv!8PE"))
	if string(second) != "Hello World!" {
		t.Fatalf("Decoded 'OwJh>Io0Tv!8PE' should be 'Hello World!', isntead got=%v", string(second))
	}
}
