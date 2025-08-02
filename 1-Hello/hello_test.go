package main 

import "testing"

func TestHello(t *testing.T) {
  t.Run("hello world", func(t *testing.T) {
    got := Hello()
    want := "Hello world"

    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  })

  t.Run("Hello with custome value string", func(t *testing.T) {
    got := HelloCustom("aji")
		want := "hello, aji"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
  })
} 
