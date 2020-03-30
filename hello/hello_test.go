package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Chris","")
		want := "Hello, Chris"
		assertCorrectMessage(t,got,want)
	})

	t.Run("saying 'Hello World' when an empty string is supplied",func(t *testing.T){
		got := Hello("","")
		want := "Hello, World"
		assertCorrectMessage(t,got,want)
	})

	t.Run("saying hellow in spanish",func(t *testing.T){
		got := Hello("Chris","Spanish")
		want := "Hola, Chris"
		assertCorrectMessage(t,got,want)
	})

	t.Run("saying hellow in french",func(t *testing.T){
		got := Hello("Chris","French")
		want := "Bonjour, Chris"
		assertCorrectMessage(t,got,want)
	})

}