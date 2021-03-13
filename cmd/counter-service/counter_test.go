package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func init() {
	counterLimit = 5
	sc = NewCount()
}

// FIxME:
func TestCallDouble(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		input        int
		counterLimit int
		expected     int
	}{
		{
			"test1",
			1,
			1,
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counterLimit = tt.counterLimit
			d := callDouble(tt.input)
			// t.Log("output: ", counterLimit)
			require.Equal(t, tt.expected, d)
		})
	}
}
