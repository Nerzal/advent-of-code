package game

import "unicode"

type Service struct {
	input []string
}

func NewService(input []string) *Service {
	return &Service{
		input: input,
	}
}

func (s *Service) Convert() Game {
	result := Game{
		Data: [][]string{},
	}

	for _, row := range s.input {
		rowData := []string{}

		for _, character := range row {
			if unicode.IsSpace(character) {
				continue
			}

			if unicode.IsDigit(character) {
				rowData = append(rowData, string(character))
				continue
			}

			if string(character) == "." {
				rowData = append(rowData, ".")
				continue
			}

			rowData = append(rowData, "#")
		}

		result.Data = append(result.Data, rowData)
	}

	return result
}
