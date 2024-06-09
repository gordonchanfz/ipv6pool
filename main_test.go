package main

import (
	"fmt"
	"testing"
)

func TestIP(t *testing.T) {
	randIP, err := randomIPV6FromSubnet("2001:470:c:5b::/64")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(randIP)
}
