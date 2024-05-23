package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(str string) string {
	// 创建一个 SHA1 散列对象
	hash := sha256.New()
	// 写入要加密的数据
	hash.Write([]byte(str))
	// 计算散列值
	hashed := hash.Sum(nil)
	// 将散列值转换为十六进制字符串
	return hex.EncodeToString(hashed)
}
