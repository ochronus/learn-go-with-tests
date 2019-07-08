package iteration

const defaultIterNum = 5

// Repeat returns a string repeated 5 times
func Repeat(core string, optiternum ...int) string {
	result := ""
	iternum := defaultIterNum

	if len(optiternum) > 0 {
		iternum = optiternum[0]
	}
	if iternum == 0 {
		iternum = defaultIterNum
	}

	for i := 0; i < iternum; i++ {
		result += core
	}
	return result
}
