package main

import "testing"

func TestLuasTabung(t *testing.T) {
	if LuasTabung(10, 10) == 1256 {
		t.Log("Sukses")
	}
	if LuasTabung(20, 4) == 602.88 {
		t.Log("Sukses")
	}
	if LuasTabung(10, 8) == 904.32 {
		t.Log("Sukses")
	}
	if LuasTabung(9, 9) == 1017.36 {
		t.Log("Sukses")
	}
	if LuasTabung(9, 10) == 1193.2 {
		t.Log("Sukses")
	}
}
