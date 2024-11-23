package helpers

func LongestString(data []string) string {
	longestItem := data[0]
	for i := 1; i < len(data); i++ {
		if len(data[i]) > len(longestItem) {
			longestItem = data[i]
		}
	}

	return longestItem
}

func ExtractColumn(data [][]string, columnIndex int) []string {
	columnItems := make([]string, len(data))
	for i := 0; i < len(data); i++ {
		columnItems[i] = data[i][columnIndex]
	}
	return columnItems
}

func LongestStringInColumn(data [][]string, columnIndex int) string {
	columnItems := ExtractColumn(data, columnIndex)
	return LongestString(columnItems)
}
