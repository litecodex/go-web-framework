package service

type PasswordEncoder interface {
	Encode(rawPassword string) string
	Matches(rawPassword, encodedPassword string) bool
}
