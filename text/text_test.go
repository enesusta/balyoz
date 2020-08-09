package text

import (
	"log"
	"testing"
)

func TestCapitalizeWithTurkish(t *testing.T) {

	kutahya := CapitalizeWithTurkish("kütahya")
	log.Println(kutahya)

	if kutahya != "Kütahya" {
		t.Errorf("Test failed, expected %v, go %v", "Kütahya", kutahya)
	}

	istanbul := CapitalizeWithTurkish("istanbul")
	log.Println(istanbul)

	if istanbul != "İstanbul" {
		t.Errorf("Test failed, expected %v, go %v", "İstanbul", istanbul)
	}

	isparta := CapitalizeWithTurkish("ısparta")
	log.Println(isparta)

	if isparta != "Isparta" {
		t.Errorf("Test failed, expected %v, go %v", "Isparta", isparta)
	}

}
