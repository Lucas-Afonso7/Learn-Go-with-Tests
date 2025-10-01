package main

const helloTest = "hello, "
const espanhol = "espanhol"
const helloSpanish = "Hola, "
const portugues = "portugues"
const helloPortuguese = "Ola, "
const hellofrances = "Bonjour, "
const frances = "frances"

func hello(name string, language string) string {
    if name == "" {
        name = "world"
    }
    
	return greetingprefix(language) + name
}

func greetingprefix(language string) (prefix string) {
	prefix = helloTest
	switch language {
	case espanhol:
		prefix = helloSpanish
    case portugues:
		prefix = helloPortuguese
	case frances:
		prefix = hellofrances
	}
    
	return
}