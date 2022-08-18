package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type parseResult struct {
	w, h int
	err  error
}

func TestParseResize(t *testing.T) {
	testCases := map[string]parseResult{
		"":      {0, 0, nil},
		"0":     {0, 0, nil},
		"-0":    {0, 0, nil},
		"1":     {1, 0, nil},
		"10":    {10, 0, nil},
		"0x":    {0, 0, nil},
		"-0x":   {0, 0, nil},
		"1x":    {1, 0, nil},
		"10x":   {10, 0, nil},
		"10x20": {10, 20, nil},
	}
	for input, result := range testCases {
		w, h, err := parseResize(input)
		require.NoError(t, err, fmt.Sprintf("input string: '%s'", input))
		assert.Equal(t, result.w, w, fmt.Sprintf("input string: '%s'", input))
		assert.Equal(t, result.h, h, fmt.Sprintf("input string: '%s'", input))
	}
}

func TestParseResizeInvalid(t *testing.T) {
	testCases := []string{
		" ",
		"-1",
		"-10",
		"-1x",
		"-10x",
		"-10x0",
		"-10x1",
		"-10x10",
		"10x-1",
		"10x-10",
		"-1x-1",
		"-10x-10",
	}
	for _, input := range testCases {
		_, _, err := parseResize(input)
		require.Error(t, err, fmt.Sprintf("input string: '%s'", input))
	}
}
