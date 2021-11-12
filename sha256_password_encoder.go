package go_spring_password_encoder

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
)

// SHA256PasswordEncoder is compatible to spring's StandardPasswordEncoder.
type SHA256PasswordEncoder struct {
	saltLengthBytes int
	iterations      int
}

func NewDefaultSHA256PasswordEncoder() PasswordEncoder {
	return NewSHA256PasswordEncoder(8, 1024)
}

func NewSHA256PasswordEncoder(saltLengthBytes int, iterations int) PasswordEncoder {
	encoder := SHA256PasswordEncoder{saltLengthBytes: saltLengthBytes, iterations: iterations}
	return PasswordEncoder(&encoder)
}

func (encoder *SHA256PasswordEncoder) Matches(plainPassword string, encodedPasswordHash string) (bool, error) {
	if len(plainPassword) == 0 || len(encodedPasswordHash) == 0 {
		return false, nil
	}

	hashedPasswordBytes, err := hex.DecodeString(encodedPasswordHash)
	if err != nil {
		return false, err
	}

	saltBytes := hashedPasswordBytes[0:encoder.saltLengthBytes]

	sha256Hash := sha256.New()

	hashBytes := []byte{}
	hashBytes = append(hashBytes, saltBytes...)
	plainPasswordBytes := []byte(plainPassword)
	hashBytes = append(hashBytes, plainPasswordBytes...)
	for i := 0; i < encoder.iterations; i++ {
		sha256Hash.Reset()
		sha256Hash.Write(hashBytes)
		hashBytes = sha256Hash.Sum([]byte{})
	}

	expectedHashBytes := hashedPasswordBytes[encoder.saltLengthBytes:]

	return reflect.DeepEqual(hashBytes, expectedHashBytes), nil
}
