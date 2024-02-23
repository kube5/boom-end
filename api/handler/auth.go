package handler

import (
	"fmt"
	"os"
	"strings"

	derrors "github.com/Mu-Exchange/Mu-End/common/errors"
	"github.com/Mu-Exchange/Mu-End/common/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const AccessDetails = "access_details"

func ExtractToken(c *gin.Context) string {
	bearToken := c.GetHeader("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	return verifyToken(tokenString)
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_TOKEN")), nil
	})
	if err != nil && strings.Contains(err.Error(), "expired") {
		return nil, derrors.ErrAuthValidationExpired
	}

	if err != nil {
		return nil, derrors.ErrAuthToken
	}
	return token, nil
}

func TokenValid(c *gin.Context) error {
	token, err := VerifyToken(c)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return derrors.ErrAuthToken
	}
	return nil
}

func (s *Server) ExtractTokenMetadata(c *gin.Context) (*model.AccessDetails, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, derrors.ErrAuthToken
	}
	return getAccessDetails(token, s)
}
func (s *Server) ExtractWsTokenMetadata(params string) (*model.AccessDetails, error) {
	params = strings.TrimPrefix(params, "{")
	params = strings.TrimSuffix(params, "}")
	split := strings.Split(params, "=")
	if len(split) != 2 {
		return nil, derrors.ErrAuthToken
	}
	token, err := verifyToken(split[1])
	if err != nil {
		return nil, derrors.ErrAuthToken
	}
	return getAccessDetails(token, s)
}

func (s *Server) ExtractTokenMetadataByBody(tokenString string) (*model.AccessDetails, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_TOKEN")), nil
	})
	if err != nil && strings.Contains(err.Error(), "expired") {
		return nil, derrors.ErrAuthValidationExpired
	}

	if err != nil {
		return nil, derrors.ErrAuthToken
	}
	return getAccessDetails(token, s)
}

func getAccessDetails(token *jwt.Token, s *Server) (*model.AccessDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, derrors.ErrAuthToken
		}
		userId := fmt.Sprintf("%s", claims["user_id"])
		ad := &model.AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}
		if err := s.tc.FetchAuth(ad); err != nil {
			return nil, derrors.ErrAuthToken
		}
		return ad, nil
	}
	return nil, derrors.ErrAuthToken
}

func (s *Server) TokenAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token, err := s.ExtractTokenMetadata(c)
		if err != nil {
			s.failResponseWithErr(c, err)
			c.Abort()
			return
		}
		c.Set(AccessDetails, token)
		c.Next()
	}
}
