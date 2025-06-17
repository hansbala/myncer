package core

type Prompt interface {
	// Render it as a string.
	Render() string
}
