package handler

import (
	"context"
	"fmt"
	"time"

	derrors "github.com/Mu-Exchange/Mu-End/common/errors"
	"github.com/Mu-Exchange/Mu-End/common/model"
	"github.com/Mu-Exchange/Mu-End/common/redis"
	"github.com/Mu-Exchange/Mu-End/common/repo"
)

var ctx = context.TODO()

func NewRedis(cfg repo.CommonComponents) (*redis.Cache, error) {
	return redis.New(cfg.Cfg.Redis)
}

type TokenCache interface {
	FetchAuth(authD *model.AccessDetails) error
	TGFetchAuth(authD *model.AccessDetails) error
	NonceVerify(nonce string) error
}

type Token struct {
	*redis.Cache
}

func (t *Token) TGFetchAuth(authD *model.AccessDetails) error {
	_, err := t.Get(ctx, tgauthKey(authD.UserId, authD.AccessUuid)).Result()
	return err
}

func NewTokenCache(cache *redis.Cache) (TokenCache, error) {
	return &Token{cache}, nil
}

func (t *Token) FetchAuth(authD *model.AccessDetails) error {
	_, err := t.Get(ctx, authKey(authD.UserId, authD.AccessUuid)).Result()
	return err
}

func authKey(userId string, givenUuid string) string {
	return fmt.Sprintf("{Auth}Auth-%s-%s", userId, givenUuid)
}

func tgauthKey(userId string, givenUuid string) string {
	return fmt.Sprintf("{Authtg}Auth-%s-%s", userId, givenUuid)
}

func nonceKey(nonce string) string {
	return fmt.Sprintf("Nonce-%s", nonce)
}

func (t *Token) NonceVerify(nonce string) error {
	result, err := t.SetNX(ctx, nonceKey(nonce), time.Now().Unix(), 15*time.Minute).Result()
	if err != nil {
		return err
	}
	if !result {
		return derrors.ErrHeaderNonceVerify
	}
	return nil
}
