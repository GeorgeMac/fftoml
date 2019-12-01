package fftoml

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseConfigFile(t *testing.T) {
	file := strings.NewReader(`
[foo]
a = "bar"
[bar]
b = "foo"
[bar.baz]
c = 12345
[bar.foo]
d = ["a", "b", "c d", "e"]
`)

	found := map[string]string{}

	set := func(k, v string) error {
		found[k] = v
		return nil
	}

	assert.Nil(t, ParseConfigFile(file, set))

	assert.Equal(t, map[string]string{
		"foo-a":     "bar",
		"bar-b":     "foo",
		"bar-baz-c": "12345",
		"bar-foo-d": "a,b,c d,e",
	}, found)
}
