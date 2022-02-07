package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Andry")
	assert.Equal(t, "Hello Andryx", result, "Result must be 'Hello Andry'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Andry")
	assert.Equal(t, "Hello Andryx", result, "Result must be 'Hello Andry'")
	fmt.Println("TestHelloWorldAndry Done")
}

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
