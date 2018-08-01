package goutils

// Try tries fn maxTimes.
func Try(maxTimes int, fn func(int) error) (err error) {
	for i := 0; i < maxTimes; i++ {
		if err = fn(i); err == nil {
			return
		}
	}
	return
}
