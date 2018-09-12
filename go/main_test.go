package main

import (
	"fmt"
	"testing"
)

func TestNewObject(t *testing.T) {
	total := newObject(nil, objInt)
	if total == nil {
		t.Error("Memory is not allocated for new object")
	} else {
		fmt.Println(*total)
	}
}

func TestNewVm(t *testing.T) {
	n := newVM()
	if n == nil {
		t.Error("Virtual memory is not allocated.")
	}
	fmt.Println(n.stack)

}
