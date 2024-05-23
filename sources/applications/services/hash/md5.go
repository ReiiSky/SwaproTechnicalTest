package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

type MD5 struct{ key string }

func NewMD5(key string) MD5 {
	return MD5{}
}

func (m MD5) Hash(plain string) string {
	hasher := hmac.New(md5.New, []byte(m.key))
	hasher.Write([]byte(plain))

	hash := hasher.Sum(nil)
	hashed := hex.EncodeToString(hash)

	return hashed
}
