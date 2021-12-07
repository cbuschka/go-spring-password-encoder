package spring_password_encoder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Java-Code:
// @Test
// public void testMatch() {
// org.springframework.security.crypto.password.PasswordEncoder standardPasswordEncoder = new org.springframework.security.crypto.password.StandardPasswordEncoder();
// boolean matches = standardPasswordEncoder.matches("asdfasdf", "b376042b477bd18bcc8fa69c1158641c3464c964ae7fed6e1a6a4ed86f55ab0c432ef96a7d40ca85");
// assertTrue(matches);
// }
func TestSHA56MatchesIsSpringCompatible(t *testing.T) {

	encoder := NewDefaultSHA256PasswordEncoder()

	matches, err := encoder.Matches("asdfasdf", "b376042b477bd18bcc8fa69c1158641c3464c964ae7fed6e1a6a4ed86f55ab0c432ef96a7d40ca85")
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.True(t, matches)
}

func TestSHA56DoesNotMatchWrongPassword(t *testing.T) {

	encoder := NewDefaultSHA256PasswordEncoder()

	matches, err := encoder.Matches("thewrongone", "b376042b477bd18bcc8fa69c1158641c3464c964ae7fed6e1a6a4ed86f55ab0c432ef96a7d40ca85")
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.False(t, matches)
}

func TestSHA56Encode(t *testing.T) {

	encoder := NewDefaultSHA256PasswordEncoder()

	encodedPasswordHash, err := encoder.Encode("asdfasdf")
	if err != nil {
		t.Fatal(err)
		return
	}

	matches, err := encoder.Matches("asdfasdf", encodedPasswordHash)
	if err != nil {
		t.Fatal(err)
		return
	}
	assert.True(t, matches)

	matches, err = encoder.Matches("thewrongone", encodedPasswordHash)
	if err != nil {
		t.Fatal(err)
		return
	}
	assert.False(t, matches)

}
