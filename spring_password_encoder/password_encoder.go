package spring_password_encoder

type PasswordEncoder interface {
	Matches(plainPassword string, encodedPasswordHash string) (bool, error)
}
