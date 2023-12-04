package game

import "strconv"

type Game struct {
	Data [][]string
}

func (g Game) Solve() int {
	result := 0

	for rowIndex, row := range g.Data {

		numberString := ""
		numberHasNeighbor := false
		for columnIndex, column := range row {
			if column == "#" || column == "." {
				if numberHasNeighbor && numberString != "" {
					number, _ := strconv.Atoi(numberString)
					result += number
				}

				numberString = ""
				numberHasNeighbor = false
				continue
			}

			numberString += column

			if g.checkRight(rowIndex, columnIndex, len(row)) ||
				g.checkTopRight(rowIndex, columnIndex) ||
				g.checkBottomRight(rowIndex, columnIndex) ||
				g.checkLeft(rowIndex, columnIndex) ||
				g.checkTopLeft(rowIndex, columnIndex) ||
				g.checkBottomLeft(rowIndex, columnIndex) ||
				g.checkBottom(rowIndex, columnIndex) ||
				g.checkTop(rowIndex, columnIndex) {
				numberHasNeighbor = true
			}
		}
	}

	return result
}

func (g Game) checkRight(rowIndex, columnIndex, lenRow int) bool {
	if columnIndex+1 < lenRow {
		// Rechter Nachbar
		if g.Data[rowIndex][columnIndex+1] == "#" {
			return true
		}
	}

	return false
}

func (g Game) checkTopRight(rowIndex, columnIndex int) bool {
	if rowIndex-1 < 0 {
		return false
	}

	if columnIndex+1 < len(g.Data[rowIndex-1]) {
		if g.Data[rowIndex-1][columnIndex+1] == "#" {
			return true
		}
	}

	return false
}

func (g Game) checkBottomRight(rowIndex, columnIndex int) bool {
	if rowIndex+1 >= len(g.Data) {
		return false
	}

	rowLen := len(g.Data[rowIndex+1])

	if columnIndex+1 < rowLen {
		if g.Data[rowIndex+1][columnIndex+1] == "#" {
			return true
		}
	}

	return false
}

func (g Game) checkLeft(rowIndex, columnIndex int) bool {
	if columnIndex-1 >= 0 {
		if g.Data[rowIndex][columnIndex-1] == "#" {
			return true
		}
	}

	return false
}

func (g Game) checkTopLeft(rowIndex, columnIndex int) bool {
	if rowIndex-1 < 0 {
		return false
	}

	if columnIndex-1 < 0 {
		return false
	}

	if g.Data[rowIndex-1][columnIndex-1] == "#" {
		return true
	}

	return false
}

func (g Game) checkBottomLeft(rowIndex, columnIndex int) bool {
	if rowIndex+1 >= len(g.Data) {
		return false
	}

	if columnIndex-1 < 0 {
		return false
	}

	if g.Data[rowIndex+1][columnIndex-1] == "#" {
		return true
	}

	return false
}

func (g Game) checkTop(rowIndex, columnIndex int) bool {
	if rowIndex-1 < 0 {
		return false
	}

	if g.Data[rowIndex-1][columnIndex] == "#" {
		return true
	}

	return false
}

func (g Game) checkBottom(rowIndex, columnIndex int) bool {
	if rowIndex+1 >= len(g.Data) {
		return false
	}

	if g.Data[rowIndex+1][columnIndex] == "#" {
		return true
	}

	return false
}
