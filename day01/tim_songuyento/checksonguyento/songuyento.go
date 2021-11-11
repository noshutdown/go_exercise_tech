package songuyento

func CheckPrime(n int) []int {

	prime := []int{}
	for n2 := 0; n2 <= n; n2++ {
		for i := 2; i < n2; i++ {
			if n2%i == 0 {
				break
			} else if i == n2-1 {
				prime = append(prime, n2)
			}
		}
	}
	return prime
}
