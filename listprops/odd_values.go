package listprops

// OddValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der ungeraden Zahlen in der Liste.
func OddValues(list []int) int {
	var x int = 0
	for i := 0; i < len(list); i++ {
		value := list[i]
		if value%2 > 0 {
			x++
		}
	}
	return x
}
