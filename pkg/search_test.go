package carve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	v, err := getVersion("../.git")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "v0.0.1", v)
}
