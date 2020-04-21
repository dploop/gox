package json

import (
	"bytes"
	"encoding/json"
)

func UnsafeMarshal(v interface{}) []byte {
	buf, _ := json.Marshal(v)
	return buf
}

func UnsafeMarshalIndent(v interface{}, prefix, indent string) []byte {
	buf, _ := json.MarshalIndent(v, prefix, indent)
	return buf
}

func MarshalNoEscape(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func MarshalIndentNoEscape(v interface{}, prefix, indent string) ([]byte, error) {
	src, err := MarshalNoEscape(v)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, src, prefix, indent)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func UnsafeMarshalNoEscape(v interface{}) []byte {
	buf, _ := MarshalNoEscape(v)
	return buf
}

func UnsafeMarshalIndentNoEscape(v interface{}, prefix, indent string) []byte {
	buf, _ := MarshalIndentNoEscape(v, prefix, indent)
	return buf
}
