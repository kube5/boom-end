package cache

import (
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

type TgUserCache interface {
	SetMetamaskLoginNonce(publicAddress, nonce string) error
	GetMetamaskLoginNonce(publicAddress string) (string, error)
	SetTgUser(user *dto.TgUser) error
	GetTgUser(userId string) (*dto.TgUser, error)
	SetAddressTgUser(user *dto.TgUser) error
	GetAddressTgUser(wallet string) (*dto.TgUser, error)
	DeleteTgUser(userId string) error
	CreateAuth(userid string, td *model.TokenDetails) error
	DeleteAuth(userid string, givenUuid string) (int64, error)
	DeleteTgUserAuth(userId string) error
	SetLoginTime(userId string) error
	GetLoginTime(userId string) (int64, error)

	SetInviteCodeSeq(seq int64) error
	GetInviteCodeSeq() (int64, error)
	InitInviteCodeSeq() error
}

type TgUser struct {
	*redis.Cache
}

const TGServicePrefix = "{TGSUSER}"

func (u *TgUser) SetInviteCodeSeq(seq int64) error {
	return u.Set(ctx, InviteCodeSeqKey, seq, 0).Err()
}

func (u *TgUser) InitInviteCodeSeq() error {
	return u.SetNX(ctx, InviteCodeSeqKey, 0, 0).Err()
}
func (u *TgUser) GetInviteCodeSeq() (int64, error) {
	return u.Get(ctx, InviteCodeSeqKey).Int64()
}

const InviteCodeSeqKey = TGServicePrefix + "InviteCodeSeq"

func (u *TgUser) Start(sd fx.Shutdowner) error {
	return nil
}

func (u *TgUser) Stop(sd fx.Shutdowner) error {
	return nil
}

func NewTgUserCache(cache *redis.Cache) (TgUserCache, error) {
	return &TgUser{cache}, nil
}

func HIdtgKey(id string) string {
	return fmt.Sprintf("%s-%s", TGServicePrefix, id)
}

func usertgKey() string {
	return fmt.Sprintf("%s-USER", TGServicePrefix)
}

func addressTgUserKey() string {
	return fmt.Sprintf("%s-ADDR-USER", TGServicePrefix)
}

func metamasktgLoginNonceKey(publicAddress string) string {
	return fmt.Sprintf("%s-METAMASK-LOGIN-NONCE-%s", TGServicePrefix, publicAddress)
}

func tgauthKey(userId string, givenUuid string) string {
	return fmt.Sprintf("{Authtg}Auth-%s-%s", userId, givenUuid)
}

func tgauthPrefixKey(userId string) string {
	return fmt.Sprintf("{Authtg}Auth-%s*", userId)
}
func tgUsers() string {
	return fmt.Sprintf("{TGSUSER}-USER-TG")
}

func (u *TgUser) CreateAuth(userid string, td *model.TokenDetails) error {
	at := time.Unix(int64(td.AtExpires), 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(int64(td.RtExpires), 0)
	now := time.Now()

	errAccess := u.Set(ctx, tgauthKey(userid, td.AccessUuid), userid, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := u.Set(ctx, tgauthKey(userid, td.RefreshUuid), userid, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func (u *TgUser) DeleteAuth(userid, givenUuid string) (int64, error) {
	deleted, err := u.Del(ctx, tgauthKey(userid, givenUuid)).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func (u *TgUser) DeleteTgUserAuth(userId string) error {
	keys, cursor, err := u.Scan(ctx, 0, tgauthPrefixKey(userId), 1000).Result()
	if err != nil {
		return err
	}
	for cursor != 0 {
		var appendKeys []string
		appendKeys, cursor, err = u.Scan(ctx, cursor, tgauthPrefixKey(userId), 1000).Result()
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

func (u *TgUser) SetTgUser(user *dto.TgUser) error {
	if !strings.EqualFold("", user.Wallet) {
		if err := u.SetAddressTgUser(user); err != nil {
			return err
		}
	}
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := u.ZAdd(ctx, tgUsers(), redis_client.Z{Member: user.Wallet, Score: float64(0)}).Err(); err != nil {
		return err
	}
	return u.HSet(ctx, usertgKey(), HIdtgKey(user.UUID), data).Err()
}

func (u *TgUser) DeleteTgUser(userId string) error {
	return u.HDel(ctx, usertgKey(), HIdtgKey(userId)).Err()
}

func (u *TgUser) SetAddressTgUser(user *dto.TgUser) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return u.HSet(ctx, addressTgUserKey(), HIdtgKey(user.Wallet), data).Err()
}

func (u *TgUser) GetAddressTgUser(wallet string) (*dto.TgUser, error) {
	user := &dto.TgUser{}
	val, err := u.HGet(ctx, addressTgUserKey(), HIdtgKey(wallet)).Result()
	if err == redis_client.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(val), user)
	return user, err
}

func (u *TgUser) GetTgUser(userId string) (*dto.TgUser, error) {
	user := &dto.TgUser{}
	val, err := u.HGet(ctx, usertgKey(), HIdtgKey(userId)).Result()
	if err == redis_client.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(val), user)
	return user, err
}

func (u *TgUser) SetLoginTime(userId string) error {
	lastLoginTime, _ := u.Get(ctx, userThisLoginTimeKey(userId)).Int64()
	if lastLoginTime > 0 {
		if err := u.Set(ctx, userLastLoginTimeKey(userId), lastLoginTime, 0).Err(); err != nil {
			return err
		}
	}
	return u.Set(ctx, userThisLoginTimeKey(userId), time.Now().UnixMilli(), 0).Err()
}

func (u *TgUser) GetLoginTime(userId string) (int64, error) {
	return u.Get(ctx, userThisLoginTimeKey(userId)).Int64()
}

func (u *TgUser) GetMetamaskLoginNonce(publicAddress string) (string, error) {
	result, err := u.Get(ctx, metamasktgLoginNonceKey(publicAddress)).Result()
	if err == redis_client.Nil {
		return "", nil
	}
	if err != nil {
		return "", nil
	}
	return result, nil
}
func (u *TgUser) SetMetamaskLoginNonce(publicAddress, nonce string) error {
	return u.Set(ctx, metamasktgLoginNonceKey(publicAddress), nonce, time.Hour*24).Err()
}
