package main

func Hello(language, name string) string {
	if name == "" {
		name = "World"
	}
	var greeting string
	switch language {
	case "spanish":
		greeting = "Hola"
	case "french":
		greeting = "Bonjour"
	default:
		greeting = "Hello"
	}
	return greeting + " " + name
}
