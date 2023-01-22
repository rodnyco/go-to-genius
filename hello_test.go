package gotogenius

import (
	"testing"
)

func TestHelloReturner(t *testing.T) {
	t.Log("Hello function can return Hello string")
	{
		name := "Rodny"
		res := Hello(name)
		ok := res == "Hi, Rodny. Welcome!"
		if !ok {
			t.Fatal("faild")
		}
	}
}
