package iteration

// Repeat gets string at returns it 6 times
func Repeat(a string, repeat int) (result string) {
	for i := 0; i < repeat; i++ {
		result += a
	}
	return result
}
