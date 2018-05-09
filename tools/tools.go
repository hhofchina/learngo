package tools

import (
	"bytes"
	"encoding/gob"
)

// 任意对象转[]byte.
func ToBytes(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 从[]byte转对象.
func FromBytes(b []byte, v interface{}) error {
	var buf bytes.Buffer
	buf.Write(b)
	if err := gob.NewDecoder(&buf).Decode(v); err != nil {
		return err
	}
	return nil
}
