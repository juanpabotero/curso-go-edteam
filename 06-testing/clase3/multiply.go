package clase3

func multiply(a, b int) int {
	return a * b
}

// funcion exportada para testear desde otro paquete
func Multiply(a, b int) int {
	return multiply(a, b)
}
