package MD5Cryt

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(params []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(params)
	return hex.EncodeToString(md5Ctx.Sum(nil))
}
