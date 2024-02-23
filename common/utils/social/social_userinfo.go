package social

import (
	"time"
)

type BaseInfoSocial struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	ProfileImageURL string `json:"profile_image_url"`
}

type DiscordInfo struct {
	Id               string      `json:"id"`
	Username         string      `json:"username"`
	DisplayName      interface{} `json:"display_name"`
	Avatar           interface{} `json:"avatar"`
	AvatarDecoration interface{} `json:"avatar_decoration"`
	Discriminator    string      `json:"discriminator"`
	PublicFlags      int         `json:"public_flags"`
	Flags            int         `json:"flags"`
	Banner           interface{} `json:"banner"`
	BannerColor      interface{} `json:"banner_color"`
	AccentColor      interface{} `json:"accent_color"`
	Locale           string      `json:"locale"`
	MfaEnabled       bool        `json:"mfa_enabled"`
	PremiumType      int         `json:"premium_type"`
}

type TwitterInfo struct {
	Data struct {
		Id              string `json:"id"`
		Name            string `json:"name"`
		Username        string `json:"username"`
		ProfileImageURL string `json:"profile_image_url"`
	} `json:"data"`
}

type Tweet struct {
	Data struct {
		Lang     string `json:"lang"`
		Entities struct {
			Mentions []struct {
				Start    int    `json:"start"`
				End      int    `json:"end"`
				Username string `json:"username"`
				Id       string `json:"id"`
			} `json:"mentions"`
			Urls []struct {
				Start       int    `json:"start"`
				End         int    `json:"end"`
				Url         string `json:"url"`
				ExpandedUrl string `json:"expanded_url"`
				DisplayUrl  string `json:"display_url"`
			} `json:"urls"`
			Hashtags []struct {
				Start int    `json:"start"`
				End   int    `json:"end"`
				Tag   string `json:"tag"`
			} `json:"hashtags"`
		} `json:"entities"`
		ReferencedTweets []struct {
			Type string `json:"type"`
			Id   string `json:"id"`
		} `json:"referenced_tweets"`
		PossiblySensitive   bool      `json:"possibly_sensitive"`
		CreatedAt           time.Time `json:"created_at"`
		EditHistoryTweetIds []string  `json:"edit_history_tweet_ids"`
		Id                  string    `json:"id"`
		AuthorId            string    `json:"author_id"`
		Text                string    `json:"text"`
	} `json:"data"`
}

//{
//    "data": {
//        "lang": "zh",
//        "entities": {
//            "mentions": [
//                {
//                    "start": 170,
//                    "end": 179,
//                    "username": "gaetbout",
//                    "id": "1075876692135239680"
//                }
//            ],
//            "urls": [
//                {
//                    "start": 180,
//                    "end": 203,
//                    "url": "https://t.co/RukNgcZIaB",
//                    "expanded_url": "https://twitter.com/gaetbout/status/1742100656356225290?s",
//                    "display_url": "twitter.com/gaetbout/statu…"
//                }
//            ],
//            "hashtags": [
//                {
//                    "start": 3,
//                    "end": 15,
//                    "tag": "StarknetDev"
//                },
//                {
//                    "start": 22,
//                    "end": 32,
//                    "tag": "CairoLang"
//                }
//            ]
//        },
//        "referenced_tweets": [
//            {
//                "type": "quoted",
//                "id": "1742100656356225290"
//            }
//        ],
//        "possibly_sensitive": false,
//        "created_at": "2024-01-03T03:49:11.000Z",
//        "edit_history_tweet_ids": [
//            "1742392637985071190"
//        ],
//        "id": "1742392637985071190",
//        "author_id": "1507010594758483982",
//        "text": "⚡️ #StarknetDev 🤖\n\n速览 #CairoLang 每周进展更新\n\n✅ 匹配 boolean simplified\n✅ 新账户属性：#[Starknet::contract(account)]\n✅ 支持更多负 impls\n✅ 现支持存储 ByteArrays\n✅ 支持 while 表达式（仅支持解析器 + 语义）\n\nh/t @gaetbout\nhttps://t.co/RukNgcZIaB"
//    }
//}
