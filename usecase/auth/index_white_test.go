package auth

import (
	"testing"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type authAndTokenStrings struct {
	*Auth
	tokens []string
}
type authSuite struct {
	suite.Suite
	auth1 authAndTokenStrings
	auth2 authAndTokenStrings
}

func (s *authSuite) Test_TAuth_CanCreateJWT() {
	auth := &Auth{JwtSecret: "Hello"}
	username := "username1"
	jwtString, err := auth.getJWTString(username)
	assert.Nil(s.T(), err)
	wantJwtString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE0NDQ0Nzg0MDAsInVzZXJuYW1lIjoidXNlcm5hbWUxIn0.iblKeSQ93z2aCTGia0H3DXsRcXPoKMJJ99RqBZ4yHh8"
	assert.Equal(s.T(), wantJwtString, jwtString)
}

func (s *authSuite) Test_TAuth_CanCreateAndValidateJWT() {

	auth := &Auth{JwtSecret: "Hello"}
	username := "username1"
	jwtString, err := auth.getJWTString(username)
	assert.Nil(s.T(), err)

	token, err := auth.validateTokenString(jwtString)

	assert.Nil(s.T(), err)
	claims, ok := token.Claims.(jwt.MapClaims)

	assert.True(s.T(), ok)

	assert.Equal(s.T(), "username1", claims["username"])

}

func (s *authSuite) GIVEN_Auth1() *authSuite {
	auth := &Auth{JwtSecret: "secretauth1"}
	s.auth1.Auth = auth
	return s
}

func (s *authSuite) GIVEN_Auth2() *authSuite {
	auth := &Auth{JwtSecret: "secretauth2"}
	s.auth2.Auth = auth
	return s
}

func (s *authSuite) WHEN_Auth1_2_SecretsDiffer() *authSuite {
	assert.NotEqual(s.T(), s.auth1.JwtSecret.asByte(), s.auth2.JwtSecret.asByte(), "[WHEN_Auth1_2_SecretsDiffer] secrets should not be the same.")
	return s
}
func (s *authSuite) GIVEN_JwtCreatedByAuth1() *authSuite {

	token, err := s.auth1.Auth.getJWTString("username1")

	assert.Nil(s.T(), err)
	s.auth1.tokens = append(s.auth1.tokens, token)
	return s
}

func (s *authSuite) SHOULD_NotValidateByAuth2() {

	token, err := s.auth2.Auth.validateTokenString(s.auth1.tokens[0])

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), token)
}

func (s *authSuite) Test_ShouldInvalidate() {
	s.GIVEN_Auth1().GIVEN_Auth2().GIVEN_JwtCreatedByAuth1().WHEN_Auth1_2_SecretsDiffer().SHOULD_NotValidateByAuth2()
}

func TestHook(t *testing.T) {
	thisSuite := new(authSuite)
	suite.Run(t, thisSuite)
}
