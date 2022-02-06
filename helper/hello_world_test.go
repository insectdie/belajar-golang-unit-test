package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldAndry(t *testing.T) {
	result := HelloWorld("oke")

	if result != "Hello Andry" {
		t.Fail()
	}

	fmt.Println("TestHelloWorldAndry Done")
}

func TestHelloWorldOmpusunggu(t *testing.T) {
	result := HelloWorld("Ompusunggu")

	if result != "Hello Ompusungguz" {
		t.Fatal("Ini Gagal")
	}

	fmt.Println("TestHelloWorldOmpusunggu Done")
}
