package kata

import (
	"testing"
)

func TestDNAtoRNA(t *testing.T) {
	//DNAtoRNA("GCAT")).To(Equal("GCAU")
	if DNAtoRNA("GCAT") != "GCAU" {
		t.Error("Expected GCAU")
	}
}
