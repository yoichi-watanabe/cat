package cat

import "bytes"

// catは +=演算子を使って文字列結合
func cat(ss ...string) string {
	var r string
	for _, s := range ss {
		r += s
	}
	return r
}

// bufは bytes.Bufferを使って文字列結合
func buf(ss ...string) string {
	var b bytes.Buffer
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}
