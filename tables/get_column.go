package tables

// GetColumn erwartet ein zweidimensionales Array (Tabelle) und eine Spaltennummer.
// Liefert eine Liste mit den Werten der angegebenen Spalte.
// Falls die Zeilen unterschiedliche Längen haben, wird für fehlende Werte ein leerer String geliefert.
func GetColumn(table [][]string, col int) []string {
	result := make([]string, 0, len(table))

	for i := 0; i < len(table); i++ {
		row := table[i]

		if col < len(row) {
			result = append(result, row[col])
		} else {
			result = append(result, "")
		}
	}

	return result

}
