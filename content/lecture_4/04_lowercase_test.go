package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestIdToLowercase(t *testing.T) {
	newId := idToLowercase(xiaomiId)

	assert.Equal(t, newId, expected)
}
