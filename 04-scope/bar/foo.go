package bar

// You demos use of block scoped vars returned and exported functions
func You(name string) string {
	message := ", you are a future " + what() + " programmer!"
	return name + message
}

// This function is not exported - lowercase name
func what() string {
	return "rock star"
}
