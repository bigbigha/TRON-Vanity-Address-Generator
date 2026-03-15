package tron

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"golang.org/x/crypto/sha3"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

const (
	// TronPrefix 是TRON地址的前缀字节
	TronPrefix = 0x41 // 'A'
)

// Base58Alphabet Base58编码的字符表
const Base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// GenerateTronAddress 生成TRON地址
func GenerateTronAddress() (string, string, error) {
	// 生成私钥和公钥
	privateKey, publicKey, err := GenerateKeyPair()
	if err != nil {
		return "", "", err
	}

	// 计算公钥的Keccak-256哈希（去掉未压缩公钥的前缀字节0x04）
	// publicKey应该是65字节：0x04 + 32字节X + 32字节Y
	keccakHash := Keccak256(publicKey[1:])

	// 取后20字节，在前面加0x41
	tronAddress := make([]byte, 21)
	tronAddress[0] = TronPrefix
	copy(tronAddress[1:], keccakHash[12:])

	// Base58Check编码
	addressBase58 := Base58CheckEncode(tronAddress)

	// 转换私钥为十六进制字符串
	privateHex := hex.EncodeToString(privateKey)

	return privateHex, addressBase58, nil
}

// GenerateTronAddressFromPrivateKey 从已知私钥生成TRON地址（用于测试）
func GenerateTronAddressFromPrivateKey(privateKeyHex string) (string, string, error) {
	// 解码私钥
	privateKey, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", "", err
	}

	if len(privateKey) != 32 {
		return "", "", fmt.Errorf("invalid private key length: %d, expected 32 bytes", len(privateKey))
	}

	// 从私钥生成公钥
	privKey := secp256k1.PrivKeyFromBytes(privateKey)
	pubKey := privKey.PubKey()
	publicKey := pubKey.SerializeUncompressed()

	// 计算公钥的Keccak-256哈希
	keccakHash := Keccak256(publicKey[1:])

	// 取后20字节，在前面加0x41
	tronAddress := make([]byte, 21)
	tronAddress[0] = TronPrefix
	copy(tronAddress[1:], keccakHash[12:])

	// Base58Check编码
	addressBase58 := Base58CheckEncode(tronAddress)

	return privateKeyHex, addressBase58, nil
}

// Base58CheckEncode 实现Base58Check编码
func Base58CheckEncode(data []byte) string {
	// 计算校验和（两次SHA256）
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])
	checksum := hash2[:4]

	// 添加校验和（创建新的切片避免修改原始数据）
	fullData := make([]byte, len(data)+4)
	copy(fullData, data)
	copy(fullData[len(data):], checksum)

	// Base58编码
	return Base58Encode(fullData)
}

// Base58Encode Base58编码
func Base58Encode(data []byte) string {
	var result []byte
	x := new(big.Int).SetBytes(data)
	base := big.NewInt(int64(len(Base58Alphabet)))
	zero := big.NewInt(0)

	for x.Cmp(zero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, base, mod)
		result = append([]byte{Base58Alphabet[mod.Int64()]}, result...)
	}

	// 处理前导零
	for _, b := range data {
		if b == 0 {
			result = append([]byte{Base58Alphabet[0]}, result...)
		} else {
			break
		}
	}

	return string(result)
}

// Keccak256 计算Keccak-256哈希（使用正确的Keccak算法）
func Keccak256(data []byte) []byte {
	h := sha3.NewLegacyKeccak256() // 使用Legacy版本以兼容以太坊/TRON
	h.Write(data)
	return h.Sum(nil)
}