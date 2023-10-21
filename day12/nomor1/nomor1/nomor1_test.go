package main

import "testing"

func TestPrintNama(t *testing.T) {
	if PrintNama("Galih") == "Hello Galih! Saya golang bahasa yang sangat menyenangkan." {
		t.Log("Sukses")
	}
	if PrintNama("Galih") == "Hello Galih!Saya golang bahasa yang sangat menyenangkan." {
		t.Error("Terjadi kesalahan")
	}
	if PrintNama("Galih") == "Hello Galih ! Saya golang bahasa yang sangat menyenangkan." {
		t.Error("Terjadi kesalahan")
	}
	if PrintNama("Galih") == "HelloGalih! Saya golang bahasa yang sangat menyenangkan." {
		t.Error("Terjadi kesalahan")
	}
	if PrintNama("Galih") == "HelloGalih ! Saya golang bahasa yang sangat menyenangkan." {
		t.Error("Terjadi kesalahan")
	}
}
