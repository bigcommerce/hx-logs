package logs_test

import (
	"bytes"
	"encoding/json"
	. "github.com/bigcommerce/logs"
	. "github.com/bigcommerce/logs/testing"
	"testing"
)

var TestTags = Tags{
	{"foo", 123},
	{"bar", 456},
	{"foo", 789},
}

func TestTags_MarshalJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	expected := `{
  "foo": 789,
  "bar": 456
}
`
	err := enc.Encode(TestTags)
	Ok(t, err)
	Equals(t, expected, buf.String())
}

func TestTags_QueryEncode(t *testing.T) {
	expected := "foo=789&bar=456"
	Equals(t, expected, string(TestTags.QueryEncode()))
}
