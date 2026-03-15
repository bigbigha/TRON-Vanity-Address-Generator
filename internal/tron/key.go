package tron

import (
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

// GenerateKeyPair 生成随机的私钥和公钥（使用secp256k1曲线）
func GenerateKeyPair() ([]byte, []byte, error) {
	// 使用secp256k1生成私钥
	priv, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		return nil, nil, err
	}

	// 获取未压缩公钥：65字节，格式为 0x04 + 32字节X坐标 + 32字节Y坐标
	pubKey := priv.PubKey()
	publicKey := pubKey.SerializeUncompressed()

	// 获取私钥的32字节表示
	privateKey := priv.Serialize()

	return privateKey, publicKey, nil
}