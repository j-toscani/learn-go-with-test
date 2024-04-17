package main

import "fmt"

// Grouping constants
const (
	german             = "German"
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	germanHelloPrefix  = "Hallo, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// Publich function as it starts with uppercase letter
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// Named return value; Initialised as `zero value` (int -> 0, string -> "")
// Private function as it starts with lowercase letter
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return // prefix
}

func main() {
	fmt.Println(Hello("World", ""))
}
