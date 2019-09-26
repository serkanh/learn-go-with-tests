package main

func Hello(language, name string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = "Hola "
	case "french":
		prefix = "Bonjour "
	default:
		prefix = "Hello "
	}
	return
}
