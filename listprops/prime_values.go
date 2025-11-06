package listprops

func PrimeValues(list []int) int {
	var count int = 0

	for i := 0; i < len(list); i++ { 
		value := list[i]

		if value <= 1{
			continue
		}

		isPrime := true 
		
		for j := 2; j*j <= value; j++ {
			if value%j == 0 {
				isPrime = false
				break         
			}
		}

		if isPrime {
			count++
		}
	}
	
	return count
}