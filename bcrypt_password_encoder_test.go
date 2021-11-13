package go_spring_password_encoder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Java-Code:
// @Test
// public void testMatch() {
// org.springframework.security.crypto.password.PasswordEncoder passwordEncoder = new BCryptPasswordEncoder();
// boolean matches = standardPasswordEncoder.matches("asdfasdf", "$2a$10$G4MsM1FcvK.Q.lbDFRPAF.Pwx9dtxnqbsRf5Ue4P3cfWj4Rlygpkq");
// assertTrue(matches);
// }
func TestBCryptMatchesIsSpringCompatible(t *testing.T) {

	encoder := NewDefaultBCryptPasswordEncoder()

	matches, err := encoder.Matches("asdfasdf", "$2a$10$G4MsM1FcvK.Q.lbDFRPAF.Pwx9dtxnqbsRf5Ue4P3cfWj4Rlygpkq")
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.True(t, matches)
}

func TestBCryptDoesNotMatchWrongPassword(t *testing.T) {

	encoder := NewDefaultBCryptPasswordEncoder()

	matches, err := encoder.Matches("thewrongone", "$2a$10$G4MsM1FcvK.Q.lbDFRPAF.Pwx9dtxnqbsRf5Ue4P3cfWj4Rlygpkq")
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.False(t, matches)
}

func TestBCryptEncode(t *testing.T) {

	encoder := NewDefaultBCryptPasswordEncoder()

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
