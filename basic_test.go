package basic_test

import (
	"testing"

	"github.com/jeromelesaux/basic"
	"github.com/stretchr/testify/assert"
)

func TestParseHelloWorld(t *testing.T) {
	code := `10 print "hello world"`
	_, err := basic.Parse(code)
	assert.NoError(t, err)
}
