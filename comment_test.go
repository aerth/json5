package json5

import (
	"bytes"
	"testing"
)

const jstr = /* what an ugly block of text */ `
{
"foo": 123,
// this is a comment
"bar": "baz",
"test": [ // !!
	"works?", // does it?
	"yes"
	]
}
`

func TestUnmarshalComment(t *testing.T) {
	m := map[string]interface{}{}
	if err := Unmarshal([]byte(jstr), &m); err != nil {
		t.Error(err)
		return
	}
	t.Log(m)
}

func TestDecoderComment(t *testing.T) {
	var (
		buf = new(bytes.Buffer)
		m   = map[string]interface{}{}
	)
	buf.WriteString(jstr)
	if err := NewDecoder(buf).Decode(&m); err != nil {
		t.Error(err)
		return
	}
	t.Log(m)
}
