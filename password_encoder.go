package spring_password_encoder

// PasswordEncoder is the interface to password encoding and matching
type PasswordEncoder interface {
	// Matches tests if a plain text password matches a previously encoded password. Use this to check if passwords are valid.
	Matches(plainPassword string, encodedPasswordHash string) (bool, error)

	// Encode encrypts and encodes a plain text password in a secure way (depends on the password encoder used). Use this
	// to encrypt a password for storage and later verification via Matches.
	Encode(plainPassword string) (string, error)
}
