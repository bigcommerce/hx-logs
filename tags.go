package logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

type Tag struct {
	Name  string
	Value interface{}
}

type TagSet interface {
	Tags() []*Tag
}

type Tags []*Tag

func (t Tags) Tags() []*Tag {
	return t
}

func (t Tags) Join(other ...*Tag) Tags {
	return append(t, other...)
}

func (t Tags) Map() (m map[string]interface{}) {
	m = make(map[string]interface{}, len(t))
	for _, tag := range t {
		m[tag.Name] = tag.Value
	}
	return
}

func (t Tags) QueryEncode() []byte {
	buf := new(bytes.Buffer)
	for i, tag := range t.Unique() {
		if i > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(tag.Name))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(fmt.Sprintf("%v", tag.Value)))
	}
	return buf.Bytes()
}

func (t Tags) MarshalJSON() (encoded []byte, err error) {
	buf := new(bytes.Buffer)
	add := func(v interface{}) error {
		if b, err := json.Marshal(v); err != nil {
			return err
		} else {
			buf.Write(b)
			return nil
		}
	}
	buf.WriteByte('{')
	for i, tag := range t.Unique() {
		if i > 0 {
			buf.WriteByte(',')
		}
		if err = add(tag.Name); err != nil {
			return
		}
		buf.WriteByte(':')
		if err = add(tag.Value); err != nil {
			return
		}
	}
	buf.WriteByte('}')
	encoded = buf.Bytes()
	return
}

func (t Tags) Unique() (unique Tags) {
	positions := make(map[string]int, len(t))
	for _, tag := range t {
		if position, found := positions[tag.Name]; found {
			unique[position] = tag
		} else {
			positions[tag.Name] = len(unique)
			unique = append(unique, tag)
		}
	}
	return
}
