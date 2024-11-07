package service

type Encryptor interface {
	Encrypt(text string) string
	Decrypt(cipher string) string
}
