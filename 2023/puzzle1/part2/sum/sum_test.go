package sum_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/Nerzal/advent-of-code/2023/puzzle1/part2/sum"
	"github.com/stretchr/testify/require"
)

func TestHandleLine(t *testing.T) {
	tt := []struct {
		line string
		want int
	}{
		{
			line: "3fiveeightoneightg",
			want: 38,
		},
	}

	for _, tc := range tt {
		t.Run(tc.line, func(t *testing.T) {
			got := sum.HandleLine(tc.line)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestBuildSum(t *testing.T) {
	file, err := os.Open("../input2")
	require.NoError(t, err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := sum.BuildSum(scanner)
	require.Equal(t, 281, result)
}
