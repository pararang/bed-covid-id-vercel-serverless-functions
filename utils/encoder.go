package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

// JSONIndentFormatter ...
func JSONIndentFormatter(input interface{}) *bytes.Buffer {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	_ = enc.Encode(input)

	return buf
}

func GetMD5String(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// JSONString ...
func JSONString(input interface{}) string {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(input)

	return buf.String()
}
