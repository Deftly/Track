package leetcode

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

type Codec struct{}

func (c *Codec) Encode(strs []string) string {
	result := strings.Builder{}
	for _, s := range strs {
		result.WriteByte(byte(len(s)))
		result.WriteString(s)
	}
	return result.String()
}

func (c *Codec) Decode(strs string) []string {
	var result []string
	for i := 0; i < len(strs); i++ {
		strLen := int(strs[i])
		result = append(result, strs[i+1:i+strLen+1])
		i += strLen
	}
	return result
}

func (c *Codec) EncodeV2(strs []string) []byte {
	var buf bytes.Buffer
	for _, s := range strs {
		lenBytes := make([]byte, 2)
		lenBytes[0] = byte(utf8.RuneCountInString(s) >> 8)
		lenBytes[1] = byte(utf8.RuneCountInString(s))
		buf.Write(lenBytes)
		buf.WriteString(s)
	}
	return buf.Bytes()
}

func (c *Codec) DecodeV2(data []byte) []string {
	var strs []string
	for len(data) > 0 {
		strLen := int(data[0])<<8 | int(data[1])
		data = data[2:]
		strs = append(strs, string(data[:strLen]))
		data = data[strLen:]
	}
	return strs
}
