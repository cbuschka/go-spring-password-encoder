package go_spring_password_encoder

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"reflect"
)

// SHA256PasswordEncoder is compatible to spring's StandardPasswordEncoder.
type SHA256PasswordEncoder struct {
	saltLengthBytes int
	iterations      int
}

// NewDefaultSHA256PasswordEncoder creates a SHA256 password encoder with default params
func NewDefaultSHA256PasswordEncoder() PasswordEncoder {
	return NewSHA256PasswordEncoder(8, 1024)
}

// NewSHA256PasswordEncoder creates a SHA256 password encoder
func NewSHA256PasswordEncoder(saltLengthBytes int, iterations int) PasswordEncoder {
	encoder := SHA256PasswordEncoder{saltLengthBytes: saltLengthBytes, iterations: iterations}
	return PasswordEncoder(&encoder)
}

func (encoder *SHA256PasswordEncoder) Encode(plainPassword string) (string, error) {

	saltBytes := make([]byte, encoder.saltLengthBytes)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", err
	}

	return encoder.saltHashAndEncode(plainPassword, saltBytes)
}

// Matches checks if password is the same as an encoded password hash has been created from
func (encoder *SHA256PasswordEncoder) Matches(plainPassword string, encodedPasswordHash string) (bool, error) {
	if len(plainPassword) == 0 || len(encodedPasswordHash) == 0 {
		return false, nil
	}

	hashedPasswordBytes, err := hex.DecodeString(encodedPasswordHash)
	if err != nil {
		return false, err
	}

	saltBytes := hashedPasswordBytes[0:encoder.saltLengthBytes]
	encodedPasswordBytes := encoder.saltAndHash(plainPassword, saltBytes)

	expectedHashBytes := hashedPasswordBytes[encoder.saltLengthBytes:]
	hashBytes := encodedPasswordBytes[encoder.saltLengthBytes:]

	return reflect.DeepEqual(hashBytes, expectedHashBytes), nil
}

func (encoder *SHA256PasswordEncoder) saltHashAndEncode(plainPassword string, salt []byte) (string, error) {
	resultBytes := encoder.saltAndHash(plainPassword, salt)

	return hex.EncodeToString(resultBytes), nil
}

func (encoder *SHA256PasswordEncoder) saltAndHash(plainPassword string, salt []byte) []byte {
	sha256Hash := sha256.New()
	hashBytes := []byte{}
	hashBytes = append(hashBytes, salt...)
	plainPasswordBytes := []byte(plainPassword)
	hashBytes = append(hashBytes, plainPasswordBytes...)
	for i := 0; i < encoder.iterations; i++ {
		sha256Hash.Reset()
		sha256Hash.Write(hashBytes)
		hashBytes = sha256Hash.Sum([]byte{})
	}

	resultBytes := []byte{}
	resultBytes = append(resultBytes, salt...)
	resultBytes = append(resultBytes, hashBytes...)
	return resultBytes
}
