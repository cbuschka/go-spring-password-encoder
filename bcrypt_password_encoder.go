package go_spring_password_encoder

import (
	"golang.org/x/crypto/bcrypt"
)

// BCryptPasswordEncoder is compatible to spring's BCryptPasswordEncoder.
type BCryptPasswordEncoder struct {
	strength int
}

func NewDefaultBCryptPasswordEncoder() PasswordEncoder {
	return NewBCryptPasswordEncoder(10)
}

func NewBCryptPasswordEncoder(strength int) PasswordEncoder {
	encoder := BCryptPasswordEncoder{strength: strength}
	return PasswordEncoder(&encoder)
}

func (encoder *BCryptPasswordEncoder) Encode(plainPassword string) (string, error) {

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), encoder.strength)
	if err != nil {
		return "", err
	}
	return string(passwordBytes), nil
}

func (encoder *BCryptPasswordEncoder) Matches(plainPassword string, encodedPasswordHash string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(encodedPasswordHash), []byte(plainPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}

	return err == nil, err
}
