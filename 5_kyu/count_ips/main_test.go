package main

import "testing"

func TestIpsBewteen(t *testing.T) {
	from, to := "10.0.0.0", "10.0.0.50"
	res := IpsBetween(from, to)
	if res != 50 {
		t.Fatalf("Should be 50, got=%v", res)
	}
	from, to = "20.0.0.10", "20.0.1.0"
	res = IpsBetween(from, to)
	if res != 246 {
		t.Fatalf("Should be 246, got=%v", res)
	}
}
