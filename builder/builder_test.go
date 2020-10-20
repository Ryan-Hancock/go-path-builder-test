package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArray(t *testing.T) {
	var p PathParser
	p.Separator = "/"

	assert.Equal(t, "abc", p.ParseArray([]interface{}{"abc"}))
	assert.Equal(t, "abc", p.ParseArray([]interface{}{"", "abc"}))
	assert.Equal(t, "abc/d", p.ParseArray([]interface{}{"abc", "d"}))
	assert.Equal(t, "/abc", p.ParseArray([]interface{}{"/abc"}))
	assert.Equal(t, "/abc/d", p.ParseArray([]interface{}{"/abc", "", "d"}))
	assert.Equal(t, "/abc/d", p.ParseArray([]interface{}{"/abc", "/d"}))
	assert.Equal(t, "/abc/d/e/", p.ParseArray([]interface{}{"/abc", "/d", "e/"}))
	assert.Equal(t, "/abc/d/e/f", p.ParseArray([]interface{}{"/abc", "/d", "e/", "f"}))
	assert.Equal(t, "/abc/d/0/f", p.ParseArray([]interface{}{"/abc", "/d", 0, "f"}))
	assert.Equal(t, "/abc/d/f", p.ParseArray([]interface{}{"/abc", "/d", nil, "f"}))
	assert.Equal(t, "/abc/d/f", p.ParseArray([]interface{}{"", "/abc", "", "/d", nil, "f"}))
	assert.Equal(t, "/abc/d/f/", p.ParseArray([]interface{}{"", "", "/abc", nil, nil, "/d", nil, "f", "", "", "/"}))
	assert.Equal(t, "abc/d//e/f", p.ParseArray([]interface{}{"abc", "/d//e/", "f"}))
	assert.Equal(t, "//cdn.xyz.com/images", p.ParseArray([]interface{}{"//cdn.xyz.com", "images"}))
}
