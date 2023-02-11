package jwt

import (
	"testing"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type authAndTokenStrings struct {
	*JWT
	tokens []String
}
type authSuite struct {
	suite.Suite
	auth1 authAndTokenStrings
	auth2 authAndTokenStrings
}

func (s *authSuite) Test_TAuth_CanCreateJWT() {
	auth := &JWT{Secret: "Hello"}
	username := "username1"
	jwtString, err := auth.GenerateJWTString(username)
	assert.Nil(s.T(), err)
	wantJwtString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE0NDQ0Nzg0MDAsInVzZXJuYW1lIjoidXNlcm5hbWUxIn0.iblKeSQ93z2aCTGia0H3DXsRcXPoKMJJ99RqBZ4yHh8"
	assert.Equal(s.T(), wantJwtString, jwtString)
}

func (s *authSuite) Test_TAuth_CanCreateAndValidateJWT() {

	auth := &JWT{Secret: "Hello"}
	username := "username1"
	jwtString, err := auth.GenerateJWTString(username)
	assert.Nil(s.T(), err)

	token, err := auth.validateTokenString(jwtString)

	assert.Nil(s.T(), err)
	claims, ok := token.Claims.(jwt.MapClaims)

	assert.True(s.T(), ok)

	assert.Equal(s.T(), "username1", claims["username"])

}

func (s *authSuite) GivenAuth1() *authSuite {
	auth := &JWT{Secret: "secretauth1"}
	s.auth1.JWT = auth
	return s
}

func (s *authSuite) GivenAuth2() *authSuite {
	auth := &JWT{Secret: "secretauth2"}
	s.auth2.JWT = auth
	return s
}

func (s *authSuite) WhenAuth12Secretsdiffer() *authSuite {
	assert.NotEqual(s.T(), s.auth1.Secret.asByte(), s.auth2.Secret.asByte(), "[WHEN_Auth1_2_SecretsDiffer] secrets should not be the same.")
	return s
}
func (s *authSuite) GivenJwtcreatedbyauth1() *authSuite {

	token, err := s.auth1.JWT.GenerateJWTString("username1")

	assert.Nil(s.T(), err)
	s.auth1.tokens = append(s.auth1.tokens, token)
	return s
}

func (s *authSuite) ShouldNotvalidatebyauth2() {

	token, err := s.auth2.JWT.validateTokenString(s.auth1.tokens[0])

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), token)
}

func (s *authSuite) Test_ShouldInvalidate() {
	s.GivenAuth1().GivenAuth2().GivenJwtcreatedbyauth1().WhenAuth12Secretsdiffer().ShouldNotvalidatebyauth2()
}

func TestHook(t *testing.T) {
	thisSuite := new(authSuite)
	suite.Run(t, thisSuite)
}
