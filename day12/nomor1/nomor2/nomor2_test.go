package main

import "testing"

func TestLuasSegitiga(t *testing.T) {
	if LuasSegitiga(5, 4) == 10 {
		t.Log("Sukses")
	}
	if LuasSegitiga(10, 5) == 25 {
		t.Log("Sukses")
	}
	if LuasSegitiga(6, 8) == 24 {
		t.Log("Sukses")
	}
	if LuasSegitiga(6, 4) == 12 {
		t.Log("Sukses")
	}
	if LuasSegitiga(5, 5) == 12.5 {
		t.Log("Sukses")
	}
}
