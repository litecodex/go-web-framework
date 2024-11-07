package rsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"
)

func MustParsePublicKey(publicKeyBase64 string) *rsa.PublicKey {
	key, err := getPublicKey(publicKeyBase64)
	if err != nil {
		panic(err)
	}
	return key
}

func MustParsePrivateKey(privateKeyBase64 string) *rsa.PrivateKey {
	key, err := getPrivateKey(privateKeyBase64)
	if err != nil {
		panic(err)
	}
	return key
}

func getPublicKey(publicKey string) (*rsa.PublicKey, error) {
	if !strings.HasPrefix(publicKey, "-----BEGIN PUBLIC KEY-----") {
		publicKey = "-----BEGIN PUBLIC KEY-----\n" + publicKey + "\n-----END PUBLIC KEY-----"
	}

	// 解码PEM格式的公钥
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}
	return rsaPub, nil
}

func RSAEncrypt(message string, publicKey string) (string, error) {
	rsaPub, err := getPublicKey(publicKey)
	if err != nil {
		return "", err
	}

	// 使用公钥加密
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPub, []byte(message))
	if err != nil {
		return "", err
	}

	// 将加密结果转换为Base64编码字符串
	encryptedString := base64.StdEncoding.EncodeToString(encryptedBytes)
	return encryptedString, nil
}

func getPrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	if !strings.HasPrefix(privateKey, "-----BEGIN RSA PRIVATE KEY-----") {
		privateKey = "-----BEGIN RSA PRIVATE KEY-----\n" + privateKey + "\n-----END RSA PRIVATE KEY-----"
	}

	// 解码PEM格式的私钥
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPrivateKey, ok := priv.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}
	return rsaPrivateKey, nil
}

func RSADecrypt(encryptedMessage string, privateKey string) (string, error) {
	priv, err := getPrivateKey(privateKey)
	if err != nil {
		return "", err
	}

	// Base64解码加密消息
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedMessage)
	if err != nil {
		return "", err
	}

	// 使用私钥解密
	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, priv, encryptedBytes)
	if err != nil {
		return "", err
	}

	return string(decryptedBytes), nil
}

// 分段加密函数
func RSAEncryptLongText(message, pubKey string) (string, error) {
	publicKey, err := getPublicKey(pubKey)
	if err != nil {
		return "", err
	}
	plaintext := []byte(message)

	keySize := publicKey.Size()
	hash := sha256.New()

	// 计算每块最大加密长度
	maxEncryptBlock := keySize - 2*hash.Size() - 2

	var encryptedBytes bytes.Buffer
	for start := 0; start < len(plaintext); start += maxEncryptBlock {
		end := start + maxEncryptBlock
		if end > len(plaintext) {
			end = len(plaintext)
		}

		encryptedBlock, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext[start:end])
		if err != nil {
			return "", err
		}

		encryptedBytes.Write(encryptedBlock)
	}

	return base64.StdEncoding.EncodeToString(encryptedBytes.Bytes()), nil
}

// 分段解密函数
func RSADecryptLongText(cipherText, privKey string) (string, error) {
	privateKey, err := getPrivateKey(privKey)
	if err != nil {
		return "", err
	}

	encryptedBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	keySize := privateKey.Size()

	// 计算每块最大解密长度
	maxDecryptBlock := keySize

	var decryptedBytes bytes.Buffer
	for start := 0; start < len(encryptedBytes); start += maxDecryptBlock {
		end := start + maxDecryptBlock
		if end > len(encryptedBytes) {
			end = len(encryptedBytes)
		}

		decryptedBlock, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedBytes[start:end])
		if err != nil {
			return "", err
		}

		decryptedBytes.Write(decryptedBlock)
	}

	return decryptedBytes.String(), nil
}
