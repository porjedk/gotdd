package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const thai = "Thai"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const thaiHelloPrefix = "Sawasdee, "
const englishPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == ""{
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case thai:
		prefix = thaiHelloPrefix
	default:
		prefix = englishPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}