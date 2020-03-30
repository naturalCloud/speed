package app

import (
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/syyongx/php2go"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"reflect"
	"speed/app/lib/log"
	"strings"
	"time"
)

var (
	Config *viper.Viper
	Redis  *redis.Client
	Db     *gorm.DB
	Log    *zap.SugaredLogger
	Log0   *zap.Logger
	//var Log1 *zap.Logger
	AppName string
	AppPath string
)

func init() {

	initLog()

	initConfig()
	initAppPath()
	initAppName()

	initDb()
	initRedis()
}

func initConfig() {
	Config = viper.New()
	Config.SetConfigType("json")
	Config.SetConfigFile(".config.json")
	err := Config.ReadInConfig()
	if err != nil {
		log.Panic(err.Error())
	}

	log.Info("config init success")

}

func initAppPath() {
	s, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	AppPath = php2go.Substr(s, 0, strings.Index(s, AppName)+len(AppName))
}

func initAppName() {
	AppName = Config.GetString("appName")
}

func initRedis() {
	var (
		host = Config.GetString("cache.redis.default.host")
		pass = Config.GetString("cache.redis.default.password")
		//idleConn, _ = Conf.Int("REDIS_MIN_IDLE")
	)

	Redis := redis.NewClient(&redis.Options{
		Addr:        host + ":6379",
		Password:    pass,             // Redis账号
		DB:          0,                // Redis库
		MaxRetries:  3,                // 最大重试次数
		IdleTimeout: 10 * time.Second, // 空闲链接超时时间
		Network:     "tcp",
	})
	pong, err := Redis.Ping().Result()
	if err == redis.Nil {
		log.Panic("Redis异常", err)
	} else if err != nil {
		log.Panic("redis异常:", err.Error())
	} else {
		log.Infof("redis init success %s ", pong)
	}

}

func initDb() {

	key := "db.mysql.default."
	var (
		db       *gorm.DB
		username = Config.GetString(key + "username")
		pass     = Config.GetString(key + "password")
		host     = Config.GetString(key + "host")
		port     = Config.GetString(key + "port")
		database = Config.GetString(key + "databaseName")
		charset  = Config.GetString(key + "charset")
		dialect  = Config.GetString("db.dialect")
	)
	dsn := username + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database + "?charset=" + charset
	db, err := gorm.Open(dialect, dsn)
	if err != nil || reflect.TypeOf(db).String() != "*gorm.DB" || db == nil {
		if err == nil || db == nil {
			log.Panicf("init DB connect failed db init false ")
			return
		}
		log.Panicf("init DB connect failed, error: %s", err.Error())
	}
	err = db.DB().Ping()
	if err != nil {
		log.Panicf("init DB connect failed, error: %s", err.Error())
	}
	db.LogMode(true)

	if Config.GetString("appEnv") == "prod" {
		db.SetLogger(&log.SqlLog{})
	}
	Db = db
	log.Info("init DB connect success")

}

func initLog() {
	Log0 = log.InitLog()
	Log = Log0.Sugar()
}
