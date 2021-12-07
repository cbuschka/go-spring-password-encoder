package spring_password_encoder

import (
	"golang.org/x/crypto/bcrypt"
)

// BCryptPasswordEncoder is compatible to spring's BCryptPasswordEncoder.
type BCryptPasswordEncoder struct {
	strength int
}

// NewDefaultBCryptPasswordEncoder creates a new bcrypt password encoder with default strength 10.
func NewDefaultBCryptPasswordEncoder() PasswordEncoder {
	return NewBCryptPasswordEncoder(10)
}

// NewBCryptPasswordEncoder creates a new bcrypt password encoder with configurable strength.
func NewBCryptPasswordEncoder(strength int) PasswordEncoder {
	encoder := BCryptPasswordEncoder{strength: strength}
	return PasswordEncoder(&encoder)
}

// Encode encrypts and encodes a plain text password according to the bcrypt standard.
func (encoder *BCryptPasswordEncoder) Encode(plainPassword string) (string, error) {

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), encoder.strength)
	if err != nil {
		return "", err
	}
	return string(passwordBytes), nil
}

// Matches tests if a plain text password matches a previously encoded password. Use this to check if passwords are valid.
func (encoder *BCryptPasswordEncoder) Matches(plainPassword string, encodedPasswordHash string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(encodedPasswordHash), []byte(plainPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}

	return err == nil, err
}
