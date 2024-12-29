package aes

import (
	"fmt"
	"testing"
)

func TestAES(t *testing.T) {
	// 示例明文和密码
	plaintext := "Hello, AES-GCM in Go and JS with PBKDF2!"
	password := "thisis32bitlongpassphraseimusing" // 任意长度密码

	// 加密
	ciphertext, err := EncryptGCM(password, plaintext)
	if err != nil {
		fmt.Println("加密错误:", err)
		return
	}
	fmt.Println("加密后的密文:", ciphertext)

	// 解密
	decryptedText, err := DecryptGCM(password, ciphertext)
	if err != nil {
		fmt.Println("解密错误:", err)
		return
	}
	fmt.Println("解密后的明文:", decryptedText)
}
