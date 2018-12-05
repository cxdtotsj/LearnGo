package tempconv

//CToF C to F
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//FToC F to C
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
