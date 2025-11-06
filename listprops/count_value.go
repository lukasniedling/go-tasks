package listprops

// CountValue gibt erwartet eine Liste von Zahlen und eine Zahl v.
// Liefert die Anzahl der Vorkommen von v in der Liste.
func CountValue(list []int, v int) int {
	var x int = 0
	for  _, value := range list {
		if value == v {
			x++
		}
	}
	return x
}

