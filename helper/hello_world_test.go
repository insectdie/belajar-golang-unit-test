package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubTest(t *testing.T) {
	t.Run("Andry", func(t *testing.T) {
		result := HelloWorld("Andry")
		assert.Equal(t, "Hello Andryx", result, "Result must be 'Hello Andry'")
	})

	t.Run("Ompusunggu", func(t *testing.T) {
		result := HelloWorld("Andry")
		assert.Equal(t, "Hello Andryx", result, "Result must be 'Hello Andry'")
	})
}

func TestMain(m *testing.M) {
	//berfore
	fmt.Println("Before Unit Test")
	m.Run()
	fmt.Println("After Unit Test")
	//after
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run in Windows")
	}

	result := HelloWorld("Andry")
	assert.Equal(t, "Hello Andryx", result, "Result must be 'Hello Andry'")
}

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
