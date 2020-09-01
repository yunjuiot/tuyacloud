package tuyacloud

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Timestamp returns timestamp format for randomize.
func Timestamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
}

// HmacSha256 for HMAC-SHA-256 sign.
func HmacSha256(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	_, _ = h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
