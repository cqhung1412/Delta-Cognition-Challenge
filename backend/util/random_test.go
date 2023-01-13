package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	number := RandomInt(1, 5)
	require.NotEmpty(t, number)
	require.NotZero(t, number)
	require.GreaterOrEqual(t, number, int64(0))
	require.LessOrEqual(t, number, int64(5))
}
