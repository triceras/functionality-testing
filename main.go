package main

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const french = "French"
const spanishHellloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHellloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
