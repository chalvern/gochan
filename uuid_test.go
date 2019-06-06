package gochan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefualtUUID(t *testing.T) {
	gochanUUID = 0
	for i := 1; i < 100; i++ {
		id := defualtUUID()
		assert.Equal(t, i, id)
	}
}
