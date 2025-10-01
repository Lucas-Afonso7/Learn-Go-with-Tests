package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println(hello("world", ""))
}

func TestHello(t *testing.T) {
    checkCorrectMessage := func(t *testing.T, result, expected string)  {
        t.Helper()
        if result != expected {
            t.Errorf("result '%s' expected '%s'", result, expected)
        }
    }
    t.Run("say hello to people", func(t *testing.T) {
        result := hello("Chris", "")
        expected := "hello, Chris"

        checkCorrectMessage(t, result, expected)
    })

    t.Run("says 'Hello, world' when an empty string is passed", func(t *testing.T) {
        result := hello("", "")
        expected := "hello, world"

        checkCorrectMessage(t, result, expected)
    })

    t.Run("em portugues", func(t *testing.T) { 
        result := hello("Matheus", portugues)
        expected := "Ola, Matheus"

        checkCorrectMessage(t, result, expected)
    }) 

    t.Run("in Spanish", func(t *testing.T) {
        result := hello("Marcos", espanhol)
        expected := "Hola, Marcos"

        checkCorrectMessage(t, result, expected)
    })

    t.Run("in frances", func(t *testing.T) {
        result := hello("Marcos", frances)
        expected := "Bonjour, Marcos"

        checkCorrectMessage(t, result, expected)
    })
}