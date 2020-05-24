package iteration

func Repeat(val string, repeat int) string {
	var repeated string
	for i := 0; i < repeat; i++ {
		repeated += val
	}
	return repeated
}
