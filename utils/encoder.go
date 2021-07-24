package utils

import (
	"bytes"
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
