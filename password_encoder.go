package spring_password_encoder

// PasswordEncoder is the interface to password encoding and matching
type PasswordEncoder interface {
	Matches(plainPassword string, encodedPasswordHash string) (bool, error)

	Encode(plainPassword string) (string, error)
}
