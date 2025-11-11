package reservations

// ContainsBoth erwartet eine Liste von Strings und zwei Strings s1 und s2.
// Liefert true, falls sowohl s1 als auch s2 in der Liste enthalten sind, sonst false.
func ContainsBoth(list []string, s1 string, s2 string) bool {
	foundS1 := false
	foundS2 := false

	for i := 0; i < len(list); i++ {
		if list[i] == s1 {
			foundS1 = true
		}
		if list[i] == s2 {
			foundS2 = true
		}
		if foundS1 && foundS2 {
			return true
		}
	}
	return false
}
