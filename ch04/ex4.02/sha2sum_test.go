package main

import (
	"bytes"
	"os"
	"testing"
)

func TestSha2Sum(t *testing.T) {
	tcs := []struct {
		args []string
		hash string
	}{
		{[]string{"sha2sum", "abc"}, "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad\n"},
		{[]string{"sha2sum", "-len", "384", "abc"}, "cb00753f45a35e8bb5a03d699ac65007272c32ab0eded1631a8b605a43ff5bed8086072ba1e7cc2358baeca134c825a7\n"},
		{[]string{"sha2sum", "-len", "512", "abc"}, "ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f\n"},
		{[]string{"sha2sum", "ABC"}, "b5d4045c3f466fa91fe2cc6abe79232a1a57cdf104f7a26e716e0a1e2789df78\n"},
	}

	for _, tc := range tcs {
		*bitlen = 0
		os.Args = tc.args
		stdout = new(bytes.Buffer)
		main()
		ret := stdout.(*bytes.Buffer).String()
		if ret != tc.hash {
			t.Errorf("Failed to hash args: %v, expects: %s, results: %s.", tc.args, tc.hash, ret)
		}
	}
}
