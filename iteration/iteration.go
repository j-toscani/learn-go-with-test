package iteration

const repeatCount = 5

func Repeat(character string, repeats int) (repeated string) {
	if repeats == 0 {
		repeats = repeatCount
	}

	for i := 0; i < repeats; i++ {
		repeated += character
	}

	return
}
