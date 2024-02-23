package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Mu-Exchange/Mu-End/common/dto"
	"github.com/Mu-Exchange/Mu-End/common/model"
	"github.com/Mu-Exchange/Mu-End/common/redis"
	redis_client "github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

var ctx = context.TODO()

type UserCache interface {
	SetMetamaskLoginNonce(publicAddress, nonce string) error
	GetMetamaskLoginNonce(publicAddress string) (string, error)
	SetUser(user *dto.User) error
	GetUser(userId string) (*dto.User, error)
	SetAddressUser(user *dto.User) error
	GetAddressUser(wallet string) (*dto.User, error)
	DeleteUser(userId string) error
	CreateAuth(userid string, td *model.TokenDetails) error
	DeleteAuth(userid string, givenUuid string) (int64, error)
	DeleteUserAuth(userId string) error
	SetLoginTime(userId string) error
	GetLoginTime(userId string) (int64, error)
}

type User struct {
	*redis.Cache
}

const ServicePrefix = "{SUSER}"

func (u *User) Start(sd fx.Shutdowner) error {
	return nil
}

func (u *User) Stop(sd fx.Shutdowner) error {
	return nil
}

func NewUserCache(cache *redis.Cache) (UserCache, error) {
	return &User{cache}, nil
}

func HIdKey(id string) string {
	return fmt.Sprintf("%s-%s", ServicePrefix, id)
}

func userKey() string {
	return fmt.Sprintf("%s-USER", ServicePrefix)
}

func addressUserKey() string {
	return fmt.Sprintf("%s-ADDR-USER", ServicePrefix)
}

func userLastLoginTimeKey(userId string) string {
	return fmt.Sprintf("%s-USER-LAST-LOGIN-TIME-%s", ServicePrefix, userId)
}
func userThisLoginTimeKey(userId string) string {
	return fmt.Sprintf("%s-USER-THIS-LOGIN-TIME-%s", ServicePrefix, userId)
}

func metamaskLoginNonceKey(publicAddress string) string {
	return fmt.Sprintf("%s-METAMASK-LOGIN-NONCE-%s", ServicePrefix, publicAddress)
}

func authKey(userId string, givenUuid string) string {
	return fmt.Sprintf("{Auth}Auth-%s-%s", userId, givenUuid)
}

func authPrefixKey(userId string) string {
	return fmt.Sprintf("{Auth}Auth-%s*", userId)
}

func (u *User) CreateAuth(userid string, td *model.TokenDetails) error {
	at := time.Unix(int64(td.AtExpires), 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(int64(td.RtExpires), 0)
	now := time.Now()

	errAccess := u.Set(ctx, authKey(userid, td.AccessUuid), userid, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := u.Set(ctx, authKey(userid, td.RefreshUuid), userid, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func (u *User) DeleteAuth(userid, givenUuid string) (int64, error) {
	deleted, err := u.Del(ctx, authKey(userid, givenUuid)).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func (u *User) DeleteUserAuth(userId string) error {
	keys, cursor, err := u.Scan(ctx, 0, authPrefixKey(userId), 1000).Result()
	if err != nil {
		return err
	}
	for cursor != 0 {
		var appendKeys []string
		appendKeys, cursor, err = u.Scan(ctx, cursor, authPrefixKey(userId), 1000).Result()
		if err != nil {
			return err
		}
		keys = append(keys, appendKeys...)
	}
	if len(keys) == 0 {
		return nil
	}
	return u.Del(ctx, keys...).Err()
}

func (u *User) SetUser(user *dto.User) error {
	if !strings.EqualFold("", user.Wallet) {
		if err := u.SetAddressUser(user); err != nil {
			return err
		}
	}
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return u.HSet(ctx, userKey(), HIdKey(user.UUID), data).Err()
}

func (u *User) DeleteUser(userId string) error {
	return u.HDel(ctx, userKey(), HIdKey(userId)).Err()
}

func (u *User) SetAddressUser(user *dto.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return u.HSet(ctx, addressUserKey(), HIdKey(user.Wallet), data).Err()
}

func (u *User) GetAddressUser(wallet string) (*dto.User, error) {
	user := &dto.User{}
	val, err := u.HGet(ctx, addressUserKey(), HIdKey(wallet)).Result()
	if err == redis_client.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(val), user)
	return user, err
}

func (u *User) GetUser(userId string) (*dto.User, error) {
	user := &dto.User{}
	val, err := u.HGet(ctx, userKey(), HIdKey(userId)).Result()
	if err == redis_client.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(val), user)
	return user, err
}

func (u *User) SetLoginTime(userId string) error {
	lastLoginTime, _ := u.Get(ctx, userThisLoginTimeKey(userId)).Int64()
	if lastLoginTime > 0 {
		if err := u.Set(ctx, userLastLoginTimeKey(userId), lastLoginTime, 0).Err(); err != nil {
			return err
		}
	}
	return u.Set(ctx, userThisLoginTimeKey(userId), time.Now().UnixMilli(), 0).Err()
}

func (u *User) GetLoginTime(userId string) (int64, error) {
	return u.Get(ctx, userThisLoginTimeKey(userId)).Int64()
}

func (u *User) GetMetamaskLoginNonce(publicAddress string) (string, error) {
	result, err := u.Get(ctx, metamaskLoginNonceKey(publicAddress)).Result()
	if err == redis_client.Nil {
		return "", nil
	}
	if err != nil {
		return "", nil
	}
	return result, nil
}
func (u *User) SetMetamaskLoginNonce(publicAddress, nonce string) error {
	return u.Set(ctx, metamaskLoginNonceKey(publicAddress), nonce, time.Hour*24).Err()
}
