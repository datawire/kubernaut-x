package util

// StringGenerator creates new Strings
type StringGenerator interface {

	// Generate produces a new String
	Generate() string
}
