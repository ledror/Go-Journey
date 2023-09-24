package tempconv

// converts from Celcius to Fahrenheit
func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// converts from Fahrenheit to Celcius
func FToC(f Fahrenheit) Celcius { return Celcius((f - 32) * 5 / 9) }
