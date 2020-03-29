package greetings

var greetingsString = "Hello World"

//PrintGreetings - this is a comment, happy now??
func PrintGreetings(name string) string {
	return greetingsString + "-" + name
}
