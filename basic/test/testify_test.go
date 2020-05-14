package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTestify(t *testing.T) {
	name := "Bob"
	age := 10

	assert.NotEqual(t, "bob", name)
	assert.NotEqual(t, 20, age)

	require.NotEqual(t, "bob", name)
	require.NotEqual(t, 20, age)
}
