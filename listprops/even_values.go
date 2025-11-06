package listprops

// EvenValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der geraden Zahlen in der Liste.
func EvenValues(list []int) int {
	var x int = 0
	for i := 0; i < len(list); i++ {
		value := list[i]
		if value%2 == 0 {
			x++
		}
	}
	return x
}
