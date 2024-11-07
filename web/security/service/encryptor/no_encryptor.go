package encryptor

// 不加密
type NoEncryptor struct {
}

func NewNoEncryptor() *NoEncryptor {
	return &NoEncryptor{}
}

func (n *NoEncryptor) Encrypt(text string) string {
	return text
}

func (n *NoEncryptor) Decrypt(cipher string) string {
	return cipher
}
