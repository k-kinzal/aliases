package types

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// Hasher is a hash function interface.
type Hasher func(v interface{}) string

// MD5 is hashed with MD5.
func MD5(v interface{}) string {
	var bytes []byte
	switch raw := v.(type) {
	case string:
		bytes = []byte(raw)
	default:
		out, _ := json.Marshal(raw)
		bytes = []byte(out)
	}
	hasher := md5.New()
	hasher.Write(bytes)
	return hex.EncodeToString(hasher.Sum(nil))
}

// SHA1 is hashed with SHA1.
func SHA1(v interface{}) string {
	var bytes []byte
	switch raw := v.(type) {
	case string:
		bytes = []byte(raw)
	default:
		out, _ := json.Marshal(raw)
		bytes = []byte(out)
	}
	hasher := sha1.New()
	hasher.Write(bytes)
	return hex.EncodeToString(hasher.Sum(nil))
}

// SHA256 is hashed with SHA256.
func SHA256(v interface{}) string {
	var bytes []byte
	switch raw := v.(type) {
	case string:
		bytes = []byte(raw)
	default:
		out, _ := json.Marshal(raw)
		bytes = []byte(out)
	}
	hasher := sha256.New()
	hasher.Write(bytes)
	return hex.EncodeToString(hasher.Sum(nil))
}
