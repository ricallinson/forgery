package f

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

/*
   Reverses a slice of strings.
*/
func StringSliceReverse(in []string) (out []string) {
	for i := len(in) - 1; i >= 0; i-- {
		out = append(out, in[i])
	}
	return
}

/*
   Returns a string signed with a hash.
*/
func Sign(v string, s string) string {
	hasher := sha256.New()
	hasher.Write([]byte(v + s))
	return v + "." + Encode(string(hasher.Sum(nil)))
}

/*
   Returns a verified string or and empty string.
*/
func Unsign(v string, s string) string {
	i := strings.LastIndex(v, ".")
	if i == -1 {
		return ""
	}
	val := v[:i]
	if Sign(val, s) != v {
		return ""
	}
	return val
}

/*
   Encodes a value using base64.
*/
func Encode(value string) string {
	encoded := make([]byte, base64.URLEncoding.EncodedLen(len(value)))
	base64.URLEncoding.Encode(encoded, []byte(value))
	return string(encoded)
}

/*
   Decodes a value using base64.
*/
func Decode(value string) (string, error) {
	value = value
	decoded := make([]byte, base64.URLEncoding.DecodedLen(len(value)))
	b, err := base64.URLEncoding.Decode(decoded, []byte(value))
	if err != nil {
		return "", err
	}
	return string(decoded[:b]), nil
}
