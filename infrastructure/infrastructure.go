package infrastructure

import (
	"crypto/rsa"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const (
	APPPORT    = ""
	DBHOST     = ""
	DBPORT     = ""
	DBUSER     = ""
	DBPASSWORD = ""
	DBNAME     = ""

	HTTPSWAGGER    = ""
	ROOTPATH       = ""
	RSAPUBLICPATH  = ""
	RSAPRIVATEPATH = ""

	EXTENDHOUR        = ""
	EXTENDHOURREFRESH = ""

	NANO_TO_SECOND = 1000000000
	Extend_Hour    = 72
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

	InfoLog *log.Logger
	ErrLog  *log.Logger

	db *gorm.DB

	encodeAuth *jwtauth.JWTAuth
	decodeAuth *jwtauth.JWTAuth
	privateKey *rsa.PrivateKey
	publicKey  interface{}

	// redisURL    string
	// redisClient *redis.Client

	rsaPublicPath  string
	rsaPrivatePath string

	extendHour        int
	extendHourRefresh int
)

func getStringEnvParameter(envParam string, defaultValue string) string {
	if value, ok := os.LookupEnv(envParam); ok {
		return value
	}
	return defaultValue
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env.pro")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func loadEnvParameters() {
	root, _ := os.Getwd()

	// appPort = os.Getenv("PORT")
	// if appPort == "" {
	// 	appPort = "19001"
	// }

	dbHost = getStringEnvParameter(DBHOST, goDotEnvVariable(("DBHOST")))
	dbPort = getStringEnvParameter(DBPORT, goDotEnvVariable("DBPORT"))
	dbUser = getStringEnvParameter(DBUSER, goDotEnvVariable("DBUSER"))
	dbPassword = getStringEnvParameter(DBPASSWORD, goDotEnvVariable("DBPASSWORD"))
	dbName = getStringEnvParameter(DBNAME, goDotEnvVariable("DBNAME"))
	httpSwagger = getStringEnvParameter(HTTPSWAGGER, goDotEnvVariable("HTTPSWAGGER"))

	rootPath = getStringEnvParameter(ROOTPATH, root)
	// redisURL = getStringEnvParameter(REDISURL, goDotEnvVariable("REDISURL"))
	rsaPrivatePath = getStringEnvParameter(RSAPRIVATEPATH, root+"/infrastructure/private.pem")
	rsaPublicPath = getStringEnvParameter(RSAPUBLICPATH, root+"/infrastructure/public.pem")

	extendHour, _ = strconv.Atoi(getStringEnvParameter(EXTENDHOUR, "24"))
	extendHourRefresh, _ = strconv.Atoi(getStringEnvParameter(EXTENDHOURREFRESH, "48"))
}

func init() {
	InfoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	ErrLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	var initDB bool
	flag.BoolVar(&initDB, "db", false, "allow recreate model database in postgres")
	flag.Parse()

	loadEnvParameters()

	if err := InitDatabase(initDB); err != nil {
		ErrLog.Println(err)
	}

	if err := loadAuthToken(); err != nil {
		ErrLog.Println(err)
	}

	// if err := InitRedis(); err != nil {
	// 	log.Fatal("error initialize redis: ", err)
	// }
}

// func GetEnforce() *casbin.Enforcer {
// 	db, _ := openConnection()

// 	adapter, err := gormadapter.NewAdapterByDB(db)
// 	if err != nil {
// 		ErrLog.Println(err)
// 	}
// 	enforcer, _ := casbin.NewEnforcer("./infrastructure/rbac_model.conf", adapter)
// 	return enforcer

// }

// GetDB get database instance
func GetDB() *gorm.DB {
	return db
}

// GetDBName get database name
func GetDBName() string {
	return dbName
}

// GetHTTPSwagger export link swagger
func GetHTTPSwagger() string {
	return httpSwagger
}

// GetAppPort export app port
func GetAppPort() string {
	return appPort
}

// GetRootPath export root path system
func GetRootPath() string {
	return rootPath
}

// GetRedisClient export redis client
// func GetRedisClient() *redis.Client {
// 	return redisClient
// }

func GetExtendAccessHour() int {
	return extendHour
}

func GetExtendRefreshHour() int {
	return extendHourRefresh
}

func GetEncodeAuth() *jwtauth.JWTAuth {
	return encodeAuth
}

func GetPublicKey() interface{} {
	return publicKey
}
