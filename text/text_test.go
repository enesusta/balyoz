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

	corum := CapitalizeWithTurkish("çorum")
	log.Println(corum)

	if corum != "Çorum" {
		t.Errorf("Test failed, expected %v, go %v", "Çorum", corum)
	}

}

func TestParseInt(t *testing.T) {
	t.Parallel()

	n1 := ParseInt("10")

	if n1 != 10 {
		t.Errorf("Test failed, expected %v, go %v", 10, n1)
	}

	enes := ParseInt("enes")

	t.Log(enes)

}
