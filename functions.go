package lycosa

// If is a simulation of the ternary operator
func If(condition bool, whenTrue, whenFalse interface{}) interface{} {
	if condition {
		return whenTrue
	}
	return whenFalse
}
