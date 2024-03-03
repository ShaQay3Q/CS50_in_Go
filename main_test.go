package main

import (
	_ "image/png"

	//"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_get_string(t *testing.T) {
	var str string
	str = get_string()

	require.Equal(t, "", str)
}
