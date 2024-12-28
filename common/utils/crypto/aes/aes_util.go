package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

const (
	SaltSize   = 16     // 盐值长度（字节）
	NonceSize  = 12     // IV/Nonce 长度（字节）
	KeySize    = 32     // AES-256
	PBKDF2Iter = 100000 // PBKDF2 迭代次数
)

// deriveKey 使用 PBKDF2 从密码和盐值派生密钥
func deriveKey(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, PBKDF2Iter, KeySize, sha256.New)
}

// encryptAES_GCM 加密函数，自动生成盐值和 nonce
func encryptAES_GCM(password string, plaintext string) (string, error) {
	// 生成随机盐值
	salt := make([]byte, SaltSize)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// 派生密钥
	key := deriveKey(password, salt)

	// 创建 AES 块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 创建 GCM 模式
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 生成随机 nonce
	nonce := make([]byte, NonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// 加密
	ciphertext := aesGCM.Seal(nil, nonce, []byte(plaintext), nil)

	// 拼接盐值、nonce 和密文
	combined := append(salt, nonce...)
	combined = append(combined, ciphertext...)

	// Base64 编码
	ciphertextBase64 := base64.StdEncoding.EncodeToString(combined)
	return ciphertextBase64, nil
}

// decryptAES_GCM 解密函数，根据 Base64 密文恢复明文
func decryptAES_GCM(password string, ciphertextBase64 string) (string, error) {
	// Base64 解码
	combined, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return "", err
	}

	// 检查长度
	if len(combined) < SaltSize+NonceSize {
		return "", fmt.Errorf("密文长度不足")
	}

	// 提取盐值、nonce 和密文
	salt := combined[:SaltSize]
	nonce := combined[SaltSize : SaltSize+NonceSize]
	ciphertext := combined[SaltSize+NonceSize:]

	// 派生密钥
	key := deriveKey(password, salt)

	// 创建 AES 块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 创建 GCM 模式
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 解密
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
