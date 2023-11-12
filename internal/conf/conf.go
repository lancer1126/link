package conf

import (
	"github.com/alimy/tryst/cfg"
	"log"
)

var (
	loggerSetting            *loggerConf
	loggerFileSetting        *loggerFileConf
	loggerZincSetting        *loggerZincConf
	loggerMeiliSetting       *loggerMeiliConf
	loggerOpenObserveSetting *loggerOpenObserveConf
	sentrySetting            *sentryConf
	redisSetting             *redisConf
	PyroscopeSetting         *pyroscopeConf
	DatabaseSetting          *databaseConf
	MysqlSetting             *mysqlConf
	PostgresSetting          *postgresConf
	Sqlite3Setting           *sqlite3Conf
	PprofServerSetting       *httpServerConf
	MetricsServerSetting     *httpServerConf
	WebServerSetting         *httpServerConf
	AdminServerSetting       *httpServerConf
	SpaceXServerSetting      *httpServerConf
	BotServerSetting         *httpServerConf
	LocalossServerSetting    *httpServerConf
	FrontendWebSetting       *httpServerConf
	DocsServerSetting        *httpServerConf
	MobileServerSetting      *grpcServerConf
	AppSetting               *appConf
	CacheSetting             *cacheConf
	EventManagerSetting      *eventManagerConf
	MetricManagerSetting     *metricManagerConf
	JobManagerSetting        *jobManagerConf
	CacheIndexSetting        *cacheIndexConf
	SimpleCacheIndexSetting  *simpleCacheIndexConf
	BigCacheIndexSetting     *bigCacheIndexConf
	RedisCacheIndexSetting   *redisCacheIndexConf
	SmsJuheSetting           *smsJuheConf
	AlipaySetting            *alipayConf
	TweetSearchSetting       *tweetSearchConf
	ZincSetting              *zincConf
	MeiliSetting             *meiliConf
	ObjectStorage            *objectStorageConf
	AliOSSSetting            *aliOSSConf
	COSSetting               *cosConf
	HuaweiOBSSetting         *huaweiOBSConf
	MinIOSetting             *minioConf
	S3Setting                *s3Conf
	LocalOSSSetting          *localOssConf
	JWTSetting               *jwtConf
	WebProfileSetting        *WebProfileConf
)

func Initial() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting error: %v", err)
	}
}

func setupSetting() error {
	vp, err := newViper()
	if err != nil {
		return err
	}

	ss, kv := featuresInfoFrom(vp, "Features")
	cfg.Initial(ss, kv)

	for k, v := range mappingConfig() {
		err := vp.UnmarshalKey(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func mappingConfig() map[string]any {
	objects := map[string]any{
		"App":               &AppSetting,
		"Cache":             &CacheSetting,
		"EventManager":      &EventManagerSetting,
		"MetricManager":     &MetricManagerSetting,
		"JobManager":        &JobManagerSetting,
		"PprofServer":       &PprofServerSetting,
		"MetricsServer":     &MetricsServerSetting,
		"WebServer":         &WebServerSetting,
		"AdminServer":       &AdminServerSetting,
		"SpaceXServer":      &SpaceXServerSetting,
		"BotServer":         &BotServerSetting,
		"LocalossServer":    &LocalossServerSetting,
		"FrontendWebServer": &FrontendWebSetting,
		"DocsServer":        &DocsServerSetting,
		"MobileServer":      &MobileServerSetting,
		"CacheIndex":        &CacheIndexSetting,
		"SimpleCacheIndex":  &SimpleCacheIndexSetting,
		"BigCacheIndex":     &BigCacheIndexSetting,
		"RedisCacheIndex":   &RedisCacheIndexSetting,
		"Alipay":            &AlipaySetting,
		"SmsJuhe":           &SmsJuheSetting,
		"Pyroscope":         &PyroscopeSetting,
		"Sentry":            &sentrySetting,
		"Logger":            &loggerSetting,
		"LoggerFile":        &loggerFileSetting,
		"LoggerZinc":        &loggerZincSetting,
		"LoggerMeili":       &loggerMeiliSetting,
		"LoggerOpenObserve": &loggerOpenObserveSetting,
		"Database":          &DatabaseSetting,
		"MySQL":             &MysqlSetting,
		"Postgres":          &PostgresSetting,
		"Sqlite3":           &Sqlite3Setting,
		"TweetSearch":       &TweetSearchSetting,
		"Zinc":              &ZincSetting,
		"Meili":             &MeiliSetting,
		"Redis":             &redisSetting,
		"JWT":               &JWTSetting,
		"ObjectStorage":     &ObjectStorage,
		"AliOSS":            &AliOSSSetting,
		"COS":               &COSSetting,
		"HuaweiOBS":         &HuaweiOBSSetting,
		"MinIO":             &MinIOSetting,
		"LocalOSS":          &LocalOSSSetting,
		"S3":                &S3Setting,
		"WebProfile":        &WebProfileSetting,
	}
	return objects
}
