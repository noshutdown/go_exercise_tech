package checkprime

func Checkprime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false

		} else if i == n-1 {
			return true
		}
	}
	return false
}
