package utils

import (
	"fmt"
	"net"
	"testing"
)

func TestGetFreePort(t *testing.T) {
	t.Helper()
	addr, err := net.ResolveTCPAddr("tcp", "localhost:2379")
	if err != nil {
		t.Fatal(err)
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(l.Addr().(*net.TCPAddr).Port)
	defer l.Close()
}
