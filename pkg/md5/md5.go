package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func ToMd5(oPassword string) string {
	h := md5.New()
	h.Write([]byte("hans.com"))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
