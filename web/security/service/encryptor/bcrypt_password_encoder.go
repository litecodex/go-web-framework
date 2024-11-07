package encryptor

import "golang.org/x/crypto/bcrypt"

type BCryptPasswordEncoder struct {
}

func NewBCryptPasswordEncoder() *BCryptPasswordEncoder {
	return &BCryptPasswordEncoder{}
}

func (thiz *BCryptPasswordEncoder) Encode(rawPassword string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func (thiz *BCryptPasswordEncoder) Matches(rawPassword, encodedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(rawPassword))
	return err == nil
}
