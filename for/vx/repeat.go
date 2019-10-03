package iterate

func Repeat(char string, numrepeat int) string {
	var repeated string
	for i := 0; i < numrepeat; i++ {
		repeated += char
	}
	return repeated
}
