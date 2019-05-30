package goutils

// ReverseBytes reverse byte slice.
func ReverseBytes(data []byte) []byte {
	n := len(data)
	if n == 0 {
		return nil
	}

	temp := make([]byte, n)
	for i, d := range data {
		temp[n-i-1] = d
	}
	return temp
}
