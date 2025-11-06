package tables

// Find erwartet eine Liste und einen Wert.
// Sucht den Wert in der Liste und liefert die Position.
// Liefert -1, falls der Wert nicht in der Liste vorkommt.
func Find(list []string, v string) int {
	for i := 0; i < len(list); i++ {
		if list[i] == v {
			return i
		}
	}
	return -1
}
