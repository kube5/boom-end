package config

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-micro/plugins/v4/config/source/consul"
	"github.com/spf13/viper"
	"go-micro.dev/v4/config"

	"github.com/Mu-Exchange/Mu-End/common"
	"github.com/Mu-Exchange/Mu-End/common/proto"
	ty "github.com/Mu-Exchange/Mu-End/common/utils/types"
)

const (
	ConfigKeyMySql           = "mysql"
	ConfigKeyRedis           = "redis"
	ConfigKeyBlockchains     = "blockchains"
	ConfigKeyLog             = "log"
	ConfigKeyContracts       = "contracts"
	ConfigKeyHttp            = "http"
	ConfigKeyTracer          = "tracer"
	ConfigKeyPyth            = "pyth"
	ConfigKeyTinderTrade     = "tinder_trade"
	ConfigKeyS3              = "s3"
	ConfigKeyKeys            = "keys"
	ConfigKeySocialId        = "social_id"
	ConfigKeySocialEndpoints = "social_endpoints"
	ConfigKeySocialOauth     = "social_oauth"
)

type Config struct {
	RepoRoot        string            `json:"repo_root"`
	Consul          *Consul           `json:"consul"`
	Title           string            `json:"title"`
	Log             *Log              `json:"log"`
	DB              *DB               `json:"db"`
	Redis           *Redis            `json:"redis"`
	Pyth            *Pyth             `json:"pyth"`
	S3              *S3               `json:"s3"`
	Keys            *Keys             `json:"keys"`
	Blockchains     []*Blockchain     `json:"blockchains"`
	Contracts       []*ContractConfig `json:"contracts"`
	Tracer          *Tracer           `json:"tracer"`
	HTTP            *HTTP             `json:"http"`
	TinderTrade     *TinderTrade      `json:"tinder_trade"`
	SocialId        *SocialId         `json:"social_id" mapstructure:"social_id"`
	SocialEndpoints *SocialEndpoints  `mapstructure:"social_endpoints"`
	SocialOauth     *SocialOauth      `mapstructure:"social_oauth"`
}

type Log struct {
	Level        string
	Filename     string
	MaxAge       ty.Duration `json:"max_age"`
	MaxSize      int64       `mapstructure:"-"`
	RotationTime ty.Duration `json:"rotation_time"`
}

type Tracer struct {
	Enable     bool   `mapstructure:"enable"`
	JaegerAddr string `mapstructure:"jaeger_addr"`
}

type Hystrix struct {
	Timeout                int `mapstructure:"timeout"`
	MaxConcurrentRequests  int `mapstructure:"max_concurrent_requests"`
	RequestVolumeThreshold int `mapstructure:"request_volume_threshold"`
	SleepWindow            int `mapstructure:"sleep_window"`
	ErrorPercentThreshold  int `mapstructure:"error_percent_threshold"`
}

type MQ struct {
	Addrs []string `mapstructure:"addrs"`
}

type Consul struct {
	Addrs []string `mapstructure:"addrs"`
}

type VerifyDomain struct {
	Domain          string `mapstructure:"domain"`
	InstApplyDomain string `mapstructure:"inst_apply_domain"`
}

type S3 struct {
	Region          string `mapstructure:"region"`
	Bucket          string `mapstructure:"bucket"`
	TokenBucket     string `mapstructure:"token_bucket" json:"token_bucket"`
	AccessKeyId     string `mapstructure:"access_key_id" json:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret" json:"access_key_secret"`
}
type SES struct {
	Region          string `mapstructure:"region"`
	AccessKeyId     string `mapstructure:"access_key_id" json:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret" json:"access_key_secret"`
	Sender          string `mapstructure:"sender"`
}
type AwsMsm struct {
	Region          string `mapstructure:"region"`
	AccessKeyId     string `mapstructure:"access_key_id" json:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret" json:"access_key_secret"`
	NotifyUrl       string `mapstructure:"notify_url" json:"notify_url"`
}

type DBInfo struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	IP       string `mapstructure:"ip"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"db_name" json:"db_name"`
}

type DB struct {
	MaxOpenConns    int         `mapstructure:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns    int         `mapstructure:"max_idle_conns" json:"max_idle_conns"`
	ConnMaxLifetime ty.Duration `json:"conn_max_lifetime"`
	EnableSlave     bool        `mapstructure:"enable_slave" json:"enable_slave"`
	EnableLog       bool        `mapstructure:"enable_log" json:"enable_log"`
	Master          DBInfo      `mapstructure:"master"`
	Slave           DBInfo      `mapstructure:"slave"`
}

type Redis struct {
	Password    string      `mapstructure:"password"`
	Addr        string      `mapstructure:"addr"`
	ExpiredTime ty.Duration `json:"expired_time"`
	TLS         bool        `mapstructure:"tls"`
}

type Blockchain struct {
	ChainId      uint64   `json:"chain_id" mapstructure:"chain_id"`
	RpcUrls      []string `json:"rpc_urls" mapstructure:"rpc_urls"`
	WssUrls      []string `json:"wss_urls" mapstructure:"wss_urls"`
	ConfirmBlock uint64   `json:"confirm_block" mapstructure:"confirm_block"`
	FilterStep   uint64   `json:"filter_step" mapstructure:"filter_step"`
	StartBlock   uint64   `json:"start_block" mapstructure:"start_block"`
	PrivRpcUrl   string   `json:"priv_rpc_url" mapstructure:"priv_rpc_url"`
	Multicall    string   `json:"multicall" mapstructure:"multicall"`
	DailyBlocks  uint64   `json:"daily_blocks" mapstructure:"daily_blocks"`
	RefetchDays  uint64   `json:"refetch_days" mapstructure:"refetch_days"`
}

type ES struct {
	EnableAws bool         `mapstructure:"enable_aws" json:"enable_aws"`
	AwsES     AwsESInfo    `mapstructure:"aws_es" json:"aws_es"`
	NativeES  NativeESInfo `mapstructure:"native_es" json:"native_es"`
}

type NativeESInfo struct {
	Address       []string `mapstructure:"address"`
	Username      string   `mapstructure:"username"`
	Password      string   `mapstructure:"password"`
	SkipTLSVerify bool     `mapstructure:"skip_tls_verify" json:"skip_tls_verify"`
}

type AwsESInfo struct {
	Address         []string `mapstructure:"address"`
	Region          string   `mapstructure:"region"`
	AccessKeyId     string   `mapstructure:"access_key_id"`
	AccessKeySecret string   `mapstructure:"access_key_secret"`
}

type CoinMarket struct {
	ApiKeyId string `mapstructure:"api_key_id"`
}

type Keys struct {
	Admin        string `json:"admin" mapstructure:"admin"`
	TwitterToken string `json:"twitter_token" mapstructure:"twitter_token"`
}

type HTTP struct {
	Port            int         `mapstructure:"port" json:"port"`
	MultipartMemory int64       `mapstructure:"-"`
	ReadTimeout     ty.Duration `json:"read_timeout"`
	WriteTimeout    ty.Duration `json:"write_timeout"`
	TLSEnable       bool        `mapstructure:"tls_enable" json:"tls_enable"`
	TLSCertFilePath string      `mapstructure:"tls_cert_file_path" json:"tls_cert_file_path"`
	TLSKeyFilePath  string      `mapstructure:"tls_key_file_path" json:"tls_key_file_path"`
	RPS             int         `mapstructure:"rps" json:"rps"`
}

type ContractConfig struct {
	ChainId      uint64             `json:"chain_id" mapstructure:"chain_id"`
	Address      string             `json:"address"`
	StartBlock   uint64             `json:"start_block" mapstructure:"start_block"`
	ContractType proto.ContractType `json:"contract_type" mapstructure:"contract_type"`
}

type Pyth struct {
	RpcUrl string `json:"rpc_url" mapstructure:"rpc_url"`
	WsUrl  string `json:"ws_url" mapstructure:"ws_url"`
}

type TinderTrade struct {
	Margin          uint64     `json:"margin" mapstructure:"margin"`     // 50 => 50 USDT
	Leverage        uint64     `json:"leverage" mapstructure:"leverage"` // 10 => 10x
	Tp              uint64     `json:"tp" mapstructure:"tp"`             // 9000 => 900%
	Sl              uint64     `json:"sl" mapstructure:"sl"`             // 100 => 10%
	Slippage        uint64     `json:"slippage" mapstructure:"slippage"` // 20 => 2%
	KeeperUrl       string     `json:"keeper_url" mapstructure:"keeper_url"`
	MaxOrderPerUser uint64     `json:"max_order_per_user" mapstructure:"max_order_per_user"`
	PairDesc        []PairDesc `json:"pair_desc" mapstructure:"pair_desc"`
}

type PairDesc struct {
	ChainId   uint64   `json:"chain_id" mapstructure:"chain_id"`
	PairId    uint64   `json:"pair_id" mapstructure:"pair_id"`
	LongDesc  []string `json:"long_desc" mapstructure:"long_desc"`
	ShortDesc []string `json:"short_desc" mapstructure:"short_desc"`
}

type SocialId struct {
	TwitterClientId     string `mapstructure:"twitter_client_id"`
	TwitterClientSecret string `mapstructure:"twitter_client_secret"`
	DiscordClientId     string `mapstructure:"discord_client_id"`
	DiscordClientSecret string `mapstructure:"discord_client_secret"`
}
type SocialEndpoints struct {
	TwitterAuthUrl  string `mapstructure:"twitter_auth_url"`
	DiscordAuthUrl  string `mapstructure:"discord_auth_url"`
	TwitterTokenUrl string `mapstructure:"twitter_token_url"`
	DiscordTokenUrl string `mapstructure:"discord_token_url"`
}
type SocialOauth struct {
	Twitter     string `mapstructure:"twitter"`
	Discord     string `mapstructure:"discord"`
	RedirectUri string `mapstructure:"redirect_uri"`
}

func getConsulConfig(address string) (config.Config, error) {
	consulSource := consul.NewSource(
		consul.WithAddress(address),
		consul.WithPrefix(fmt.Sprintf("/%s/config", common.AppName)),
		consul.StripPrefix(true),
	)
	conf, err := config.NewConfig()
	if err != nil {
		return conf, err
	}
	err = conf.Load(consulSource)

	return conf, err
}

func GetConfigValue(address string, key string, value interface{}) error {
	conf, err := getConsulConfig(address)
	if err != nil {
		return err
	}

	err = conf.Get(key).Scan(value)
	if err != nil {
		return err
	}

	return nil
}

func DefaultConfig(moduleName string) (*Config, error) {
	return &Config{
		Title: fmt.Sprintf("%s configuration file", moduleName),
	}, nil
}

func LoadConfig(repoRoot string, moduleName string, envPrefix string) (*Config, error) {
	viper.SetConfigFile(filepath.Join(repoRoot, moduleName+".toml"))
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix(envPrefix)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg, err := DefaultConfig(moduleName)
	if err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	logConfig := Log{}
	if err := GetConfigValue(cfg.Consul.Addrs[0], ConfigKeyLog, &logConfig); err != nil {
		return nil, err
	}
	cfg.Log = &logConfig

	dbConfig := DB{}
	if err := GetConfigValue(cfg.Consul.Addrs[0], ConfigKeyMySql, &dbConfig); err != nil {
		return nil, err
	}
	cfg.DB = &dbConfig

	redisConfig := Redis{}
	if err := GetConfigValue(cfg.Consul.Addrs[0], ConfigKeyRedis, &redisConfig); err != nil {
		return nil, err
	}
	cfg.Redis = &redisConfig

	blockchains := make([]*Blockchain, 0)
	if err := GetConfigValue(cfg.Consul.Addrs[0], ConfigKeyBlockchains, &blockchains); err != nil {
		return nil, err
	}
	cfg.Blockchains = blockchains

	contracts := make([]*ContractConfig, 0)
	if err := GetConfigValue(cfg.Consul.Addrs[0], ConfigKeyContracts, &contracts); err != nil {
		return nil, err
	}
	cfg.Contracts = contracts

	pythConfig := Pyth{}
	if err := GetConfigValue(cfg.Consul.Addrs[0], ConfigKeyPyth, &pythConfig); err != nil {
		return nil, err
	}
	cfg.Pyth = &pythConfig

	tracerConfig := Tracer{}
	if err := GetConfigValue(cfg.Consul.Addrs[0], ConfigKeyTracer, &tracerConfig); err != nil {
		return nil, err
	}
	cfg.Tracer = &tracerConfig

	httpConfig := HTTP{}
	if err := GetConfigValue(cfg.Consul.Addrs[0], ConfigKeyHttp, &httpConfig); err != nil {
		return nil, err
	}
	cfg.HTTP = &httpConfig

	s3Config := S3{}
	if err := GetConfigValue(cfg.Consul.Addrs[0], ConfigKeyS3, &s3Config); err != nil {
		return nil, err
	}
	cfg.S3 = &s3Config

	socialIdConfig := SocialId{}
	if err := GetConfigValue(cfg.Consul.Addrs[0], ConfigKeySocialId, &socialIdConfig); err != nil {
		return nil, err
	}
	cfg.SocialId = &socialIdConfig

	socialEndpointsConfig := SocialEndpoints{}
	if err := GetConfigValue(cfg.Consul.Addrs[0], ConfigKeySocialEndpoints, &socialEndpointsConfig); err != nil {
		return nil, err
	}
	cfg.SocialEndpoints = &socialEndpointsConfig

	socialOauthConfig := SocialOauth{}
	if err := GetConfigValue(cfg.Consul.Addrs[0], ConfigKeySocialOauth, &socialOauthConfig); err != nil {
		return nil, err
	}
	cfg.SocialOauth = &socialOauthConfig

	tinderTradeConfig := TinderTrade{}
	if err := GetConfigValue(cfg.Consul.Addrs[0], ConfigKeyTinderTrade, &tinderTradeConfig); err != nil {
		return nil, err
	}
	cfg.TinderTrade = &tinderTradeConfig

	cfg.RepoRoot = repoRoot

	return cfg, nil
}

func (c *Config) Bytes() ([]byte, error) {
	ret, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
