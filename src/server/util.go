package server

import (
	"crypto/sha1"
	"encoding/base32"
	"strings"
	"time"
)

func getETag(content []byte) string {
	sha := sha1.New()
	sha.Write(content)
	return strings.ToLower(base32.HexEncoding.EncodeToString(sha.Sum(nil)))
}

func formatTimeGMT(t time.Time) string {
	return t.UTC().Format(time.RFC1123[:len(time.RFC1123)-3]) + "GMT"
}
