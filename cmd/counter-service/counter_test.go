package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func init() {
	counterLimit = 5
	sc = NewCount()
}

func TestCallDouble(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		input               int
		counterLimit        int
		expectedDoubleValue int
	}{
		{
			"input with counterLimit 1",
			1,
			1,
			2,
		},
		{
			"input with counterLimit 2",
			2,
			2,
			4,
		},
		{
			"input with counterLimit 3",
			3,
			3,
			6,
		},
		{
			"input with counterLimit 4",
			4,
			4,
			8,
		},
		{
			"input with counterLimit 5",
			5,
			5,
			10,
		},
		{
			"input with counterLimit 5",
			6,
			5,
			6,
		},
		{
			"input with counterLimit 5",
			10,
			5,
			10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counterLimit = tt.counterLimit
			d := callDouble(tt.input)
			if sc.Value() > int64(tt.counterLimit) {
				require.Equal(t, tt.input, d)
			} else {
				require.Equal(t, tt.expectedDoubleValue, d)
			}
		})
	}
}

func TestDouble(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			"input 1",
			1,
			2,
		},
		{
			"input 2",
			2,
			4,
		},
		{
			"input 3",
			3,
			6,
		},
		{
			"input 5",
			5,
			10,
		},
		{
			"input 10",
			10,
			20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := double(tt.input)
			require.Equal(t, tt.expected, d)
		})
	}
}
