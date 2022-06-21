package infrastructure

import (
	"crypto/rsa"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/go-chi/jwtauth"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const (
	APPPORT    = "APP_PORT"
	DBHOST     = "DB_HOST"
	DBPORT     = "DB_PORT"
	DBUSER     = "DB_USER"
	DBPASSWORD = "DB_PASSWORD"
	DBNAME     = "DB_NAME"

	HTTPSWAGGER = "HTTP_SWAGGER"
	ROOTPATH    = "ROOT_PATH"

	PRIVATEPASSWORD = "PRIVATE_PASSWORD"
	PRIVATEPATH     = "PRIVATE_PATH"
	PUBLICPATH      = "PUBLIC_PATH"

	REDISURL = "REDIS_URL"

	EXTENDHOUR         = "EXTEND_ACCESS_HOUR"
	EXTENDACCESSMINUTE = "EXTEND_ACCESS_MINUTE"
	EXTENDREFRESHHOUR  = "EXTEND_REFRESH_HOUR"
	NANO_TO_SECOND     = 1000000000

	KEYMATCHMODEL = "KEY_MATCH_MODEL"

	MAILSERVER  = "MAIL_SERVER"
	MAILPORT    = "MAIL_PORT"
	MAILACCOUNT = "MAIL_ACCOUNT"
	MAILPASS    = "MAIL_PASS"
)

var (
	appPort    string
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string

	httpSwagger string
	rootPath    string
	staticPath  string

	InfoLog *log.Logger
	ErrLog  *log.Logger

	db         *gorm.DB
	encodeAuth *jwtauth.JWTAuth
	decodeAuth *jwtauth.JWTAuth
	privateKey *rsa.PrivateKey
	publicKey  interface{}

	redisURL    string
	redisClient *redis.Client
	enforcer    *casbin.Enforcer

	privatePassword    string
	privatePath        string
	extendAccessMinute int
	extendRefreshHour  int

	publicPath string

	extendHour int

	keyMatchModel string

	NameRefreshTokenInCookie string
	NameAccessTokenInCookie  string

	storagePath           string
	storageProductImgPath string

	mailServer   string
	mailPort     string
	mailAccount  string
	mailPassword string
)

func getStringEnvParameter(envParam string, defaultValue string) string {
	if value, ok := os.LookupEnv(envParam); ok {
		return value
	}
	return defaultValue
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func loadEnvParameters(version int, dbNameArg string, dbPwdArg string) {
	root, _ := os.Getwd()

	appPort = os.Getenv("PORT")
	if appPort == "" {
		appPort = "19001"
	}

	// appPort = getStringEnvParameter(APPPORT, goDotEnvVariable(APPPORT))

	dbPort = getStringEnvParameter(DBPORT, goDotEnvVariable(DBPORT))
	switch version {
	case 0:
		dbHost = getStringEnvParameter(DBHOST, "localhost")
		dbUser = getStringEnvParameter(DBUSER, "postgres")
		dbPassword = getStringEnvParameter(DBPASSWORD, dbPwdArg)
		dbName = getStringEnvParameter(DBNAME, dbNameArg)
		log.Println("Enviroment: LOCALHOST")
		break

	default:
		dbHost = getStringEnvParameter(DBHOST, goDotEnvVariable(DBHOST))
		dbUser = getStringEnvParameter(DBUSER, goDotEnvVariable(DBUSER))
		dbPassword = getStringEnvParameter(DBPASSWORD, goDotEnvVariable(DBPASSWORD))
		dbName = getStringEnvParameter(DBNAME, goDotEnvVariable(DBNAME))

		// dbHost = getStringEnvParameter(DBHOST, "159.65.143.187")
		// dbUser = getStringEnvParameter(DBUSER, "cdio_user")
		// dbName = getStringEnvParameter(DBNAME, "cdio")
		// dbPassword = getStringEnvParameter(DBPASSWORD, "s3#fj@dAnU")
		log.Println("Enviroment: Development Default")
	}

	privatePassword = getStringEnvParameter(PRIVATEPASSWORD, "Nhuanhthu1")
	privatePath = getStringEnvParameter(PRIVATEPATH, root+"/infrastructure/private.pem")
	publicPath = getStringEnvParameter(PUBLICPATH, root+"/infrastructure/public.pem")

	extendHour, _ = strconv.Atoi(getStringEnvParameter(EXTENDHOUR, "720"))
	extendAccessMinute, _ = strconv.Atoi(getStringEnvParameter(EXTENDACCESSMINUTE, goDotEnvVariable(EXTENDACCESSMINUTE)))
	extendRefreshHour, _ = strconv.Atoi(getStringEnvParameter(EXTENDREFRESHHOUR, goDotEnvVariable(EXTENDREFRESHHOUR)))

	keyMatchModel = getStringEnvParameter(KEYMATCHMODEL, root+"/infrastructure/keymatch_model.conf")

	httpSwagger = getStringEnvParameter(HTTPSWAGGER, goDotEnvVariable(HTTPSWAGGER))

	redisURL = getStringEnvParameter(REDISURL, goDotEnvVariable("REDIS_URL"))

	rootPath = getStringEnvParameter(ROOTPATH, root)
	staticPath = rootPath + "/static"

	NameRefreshTokenInCookie = "refreshTokenEP"
	NameAccessTokenInCookie = "accessTokenEP"

	storagePath = rootPath + "/storage/"
	// storageProductImgPath = storagePath + "products/"

}

func init() {
	InfoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	ErrLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Get version ARGS
	var version int
	flag.IntVar(&version, "v", 1, "select version dev v1 or dev v2")
	// flag.Parse()
	var dbNameArg string
	flag.StringVar(&dbNameArg, "dbname", "postgres", "database name need to connect")

	var dbPwdArg string
	flag.StringVar(&dbPwdArg, "dbpwd", "123456", "password in database need to connect")

	var initDB bool
	flag.BoolVar(&initDB, "db", false, "allow recreate model database in postgres")

	flag.Parse()
	log.Println("database version: ", version)

	loadEnvParameters(version, dbNameArg, dbPwdArg)
	if err := loadAuthToken(); err != nil {
		log.Println("error load auth token: ", err)
	}

	// if err := InitRedis(); err != nil {
	// 	log.Fatal("error initialize redis: ", err)
	// }

	if err := InitDatabase(initDB); err != nil {
		ErrLog.Println("error initialize database: ", err)
	}

	if err := InitAuthorization(); err != nil {
		log.Fatal(err)
	}
}

func GetDBName() string {
	return dbName
}

// GetDB export db
func GetDB() *gorm.DB {
	return db
}

// GetHTTPSwagger export link swagger
func GetHTTPSwagger() string {
	return httpSwagger
}

// GetAppPort export app port
func GetAppPort() string {
	return appPort
}

// GetStaticPath export static path
func GetStaticPath() string {
	return staticPath
}

// GetEncodeAuth get token auth
func GetEncodeAuth() *jwtauth.JWTAuth {
	return encodeAuth
}

// GetDecodeAuth export decode auth
func GetDecodeAuth() *jwtauth.JWTAuth {
	return decodeAuth
}

// GetExtendAccessMinute export access extend minute
func GetExtendAccessHour() int {
	return extendHour
}

// GetExtendAccessMinute export access extend minute
func GetExtendAccessMinute() int {
	return extendAccessMinute
}

// GetExtendRefreshHour export refresh extends hour
func GetExtendRefreshHour() int {
	return extendRefreshHour
}

// GetKeyMatchModel get key match model path
func GetKeyMatchModel() string {
	return keyMatchModel
}

// GetEnforcer export enforcer
func GetEnforcer() *casbin.Enforcer {
	return enforcer
}

//GetMailParam
func GetMailParam() (string, string, string, string) {
	return mailServer, mailPort, mailAccount, mailPassword
}

// GetRedisClient export redis client
func GetRedisClient() *redis.Client {
	return redisClient
}

// GetPublicKey get public key
func GetPublicKey() interface{} {
	return publicKey
}

// GetStoragePath get path of storage
func GetStoragePath() string {
	return storagePath
}

// GetStoragePath get path of storage

// GetRPFilePath get path of storage
func GetAvatarFilePath() string {
	return storageProductImgPath
}
