package conf

import (
	"bytes"
	_ "embed"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
	"time"
)

//go:embed config-default.yaml
var configBytes []byte

type pyroscopeConf struct {
	AppName   string
	Endpoint  string
	AuthToken string
	Logger    string
}

type sentryConf struct {
	Dsn              string
	Debug            bool
	AttachStacktrace bool
	TracesSampleRate float64
	AttachLogrus     bool
	AttachGin        bool
}

type loggerConf struct {
	Level string
}

type loggerFileConf struct {
	SavePath string
	FileName string
	FileExt  string
}

type loggerZincConf struct {
	Host     string
	Index    string
	User     string
	Password string
	Secure   bool
}

type loggerMeiliConf struct {
	Host         string
	Index        string
	ApiKey       string
	Secure       bool
	MaxLogBuffer int
	MinWorker    int
}

type loggerOpenObserveConf struct {
	Host         string
	Organization string
	Stream       string
	User         string
	Password     string
	Secure       bool
	MaxLogBuffer int
	MinWorker    int
}

type httpServerConf struct {
	RunMode      string
	HttpIp       string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type grpcServerConf struct {
	Host string
	Port string
}

type appConf struct {
	RunMode               string
	MaxCommentCount       int64
	MaxWhisperDaily       int64
	MaxCaptchaTimes       int
	AttachmentIncomeRate  float64
	DefaultContextTimeout time.Duration
	DefaultPageSize       int
	MaxPageSize           int
}

type cacheConf struct {
	KeyPoolSize           int
	ClientSideCacheExpire time.Duration
	UnreadMsgExpire       int64
	UserTweetsExpire      int64
	IndexTweetsExpire     int64
	MessagesExpire        int64
	IndexTrendsExpire     int64
	TweetCommentsExpire   int64
	OnlineUserExpire      int64
	UserInfoExpire        int64
	UserProfileExpire     int64
	UserRelationExpire    int64
}

type eventManagerConf struct {
	MinWorker       int
	MaxEventBuf     int
	MaxTempEventBuf int
	MaxTickCount    int
	TickWaitTime    time.Duration
}

type metricManagerConf struct {
	MinWorker       int
	MaxEventBuf     int
	MaxTempEventBuf int
	MaxTickCount    int
	TickWaitTime    time.Duration
}

type jobManagerConf struct {
	MaxOnlineInterval     string
	UpdateMetricsInterval string
}

type cacheIndexConf struct {
	MaxUpdateQPS int
	MinWorker    int
}

type simpleCacheIndexConf struct {
	MaxIndexSize       int
	CheckTickDuration  time.Duration
	ExpireTickDuration time.Duration
}

type bigCacheIndexConf struct {
	MaxIndexPage     int
	HardMaxCacheSize int
	ExpireInSecond   time.Duration
	Verbose          bool
}

type redisCacheIndexConf struct {
	ExpireInSecond time.Duration
	Verbose        bool
}

type alipayConf struct {
	AppID             string
	PrivateKey        string
	RootCertFile      string
	PublicCertFile    string
	AppPublicCertFile string
	InProduction      bool
}

type smsJuheConf struct {
	Gateway string
	Key     string
	TplID   string
	TplVal  string
}

type tweetSearchConf struct {
	MaxUpdateQPS int
	MinWorker    int
}

type zincConf struct {
	Host     string
	Index    string
	User     string
	Password string
	Secure   bool
}

type meiliConf struct {
	Host   string
	Index  string
	ApiKey string
	Secure bool
}

type databaseConf struct {
	TablePrefix string
	LogLevel    string
}

type mysqlConf struct {
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type postgresConf map[string]string

type sqlite3Conf struct {
	Path string
}

type objectStorageConf struct {
	RetainInDays int
	TempDir      string
}

type minioConf struct {
	AccessKey string
	SecretKey string
	Secure    bool
	Endpoint  string
	Bucket    string
	Domain    string
}

type s3Conf struct {
	AccessKey string
	SecretKey string
	Secure    bool
	Endpoint  string
	Bucket    string
	Domain    string
}

type aliOSSConf struct {
	AccessKeyID     string
	AccessKeySecret string
	Endpoint        string
	Bucket          string
	Domain          string
}

type cosConf struct {
	SecretID  string
	SecretKey string
	Region    string
	Bucket    string
	Domain    string
}

type huaweiOBSConf struct {
	AccessKey string
	SecretKey string
	Endpoint  string
	Bucket    string
	Domain    string
}

type localOssConf struct {
	SavePath string
	Secure   bool
	Bucket   string
	Domain   string
}

type redisConf struct {
	InitAddress      []string
	Username         string
	Password         string
	SelectDB         int
	ConnWriteTimeout time.Duration
}

type jwtConf struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type WebProfileConf struct {
	UseFriendship             bool   `json:"use_friendship"`
	EnableTrendsBar           bool   `json:"enable_trends_bar"`
	EnableWallet              bool   `json:"enable_wallet"`
	AllowTweetAttachment      bool   `json:"allow_tweet_attachment"`
	AllowTweetAttachmentPrice bool   `json:"allow_tweet_attachment_price"`
	AllowTweetVideo           bool   `json:"allow_tweet_video"`
	AllowUserRegister         bool   `json:"allow_user_register"`
	AllowPhoneBind            bool   `json:"allow_phone_bind"`
	DefaultTweetMaxLength     int    `json:"default_tweet_max_length"`
	TweetWebEllipsisSize      int    `json:"tweet_web_ellipsis_size"`
	TweetMobileEllipsisSize   int    `json:"tweet_mobile_ellipsis_size"`
	DefaultTweetVisibility    string `json:"default_tweet_visibility"`
	DefaultMsgLoopInterval    int    `json:"default_msg_loop_interval"`
	CopyrightTop              string `json:"copyright_top"`
	CopyrightLeft             string `json:"copyright_left"`
	CopyrightLeftLink         string `json:"copyright_left_link"`
	CopyrightRight            string `json:"copyright_right"`
	CopyrightRightLink        string `json:"copyright_right_link"`
}

func (s *loggerConf) logLevel() logrus.Level {
	switch strings.ToLower(s.Level) {
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn", "warning":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.ErrorLevel
	}
}

func newViper() (*viper.Viper, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath(".")
	vp.SetConfigType("yaml")
	err := vp.ReadConfig(bytes.NewReader(configBytes))
	if err != nil {
		return nil, err
	}
	// 与根目录的配置文件合并
	if err = vp.MergeInConfig(); err != nil {
		return nil, err
	}
	return vp, nil
}

func featuresInfoFrom(vp *viper.Viper, k string) (map[string][]string, map[string]string) {
	sub := vp.Sub(k)
	keys := sub.AllKeys()

	suites := make(map[string][]string)
	kv := make(map[string]string, len(keys))
	for _, key := range sub.AllKeys() {
		val := sub.Get(key)
		switch v := val.(type) {
		case string:
			kv[key] = v
		case []any:
			suites[key] = sub.GetStringSlice(key)
		}
	}
	return suites, kv
}

func timeSetting() {
	CacheSetting.ClientSideCacheExpire *= time.Second
	EventManagerSetting.TickWaitTime *= time.Second
	MetricManagerSetting.TickWaitTime *= time.Second
	JWTSetting.Expire *= time.Second
	SimpleCacheIndexSetting.CheckTickDuration *= time.Second
	SimpleCacheIndexSetting.ExpireTickDuration *= time.Second
	BigCacheIndexSetting.ExpireInSecond *= time.Second
	RedisCacheIndexSetting.ExpireInSecond *= time.Second
	redisSetting.ConnWriteTimeout *= time.Second
}
