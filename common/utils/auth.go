package utils

import (
	"os"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"

	"github.com/Mu-Exchange/Mu-End/common/model"
)

func CreateToken(userid string) (*model.TokenDetails, error) {
	td := &model.TokenDetails{}
	td.AtExpires = time.Duration(time.Now().Add(time.Hour * 24 * 7).Unix())
	accessUuid, _ := uuid.NewV4()
	td.AccessUuid = accessUuid.String()

	td.RtExpires = time.Duration(time.Now().Add(time.Hour * 24 * 14).Unix())
	refreshUuid, _ := uuid.NewV4()
	td.RefreshUuid = refreshUuid.String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_TOKEN")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["access_uuid"] = td.AccessUuid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_TOKEN")))
	if err != nil {
		return nil, err
	}
	return td, nil
}
