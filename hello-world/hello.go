package hello

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	spanishHelloPrefox = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	englishHelloPrefix = "Hello, "
)

func greetingPrefix(language string) string {
	prefix := englishHelloPrefix
	switch language {
	case spanish:
		prefix = spanishHelloPrefox
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("User", ""))
}
