package basic_test

import (
	"testing"

	"github.com/jeromelesaux/basic"
	"github.com/stretchr/testify/assert"
)

func TestParseHelloWorld(t *testing.T) {
	code := `10 print "hello world":print ABS$(35)`
	_, err := basic.Parse(code)
	assert.NoError(t, err)
}

func TestSplitLine(t *testing.T) {
	t.Run("simple line", func(t *testing.T) {
		v := basic.Split("10 print \"hello world\": print \"how are you\"")
		assert.Equal(t, 2, len(v))
		assert.False(t, v[0].IsNewLine)
		assert.False(t, v[1].IsNewLine)
	})

	t.Run("two points included in string", func(t *testing.T) {
		v := basic.Split("10 print \"hello world : \": print \"how are you\"")
		assert.Equal(t, 2, len(v))
		assert.False(t, v[0].IsNewLine)
		assert.False(t, v[1].IsNewLine)
	})
}
