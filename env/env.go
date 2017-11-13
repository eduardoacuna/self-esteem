package env

import (
	"context"
	"os"

	"github.com/eduardoacuna/self-esteem/log"
)

var Main string
var Bin string
var Address string
var Port string

var DBUser string
var DBName string
var DBHost string
var DBPort string

func Setup(ctx context.Context) {
	log.Info(ctx, "environment variables setup")
	Main = os.Getenv("SELFESTEEM_MAIN")
	Bin = os.Getenv("SELFESTEEM_BIN")
	Address = os.Getenv("SELFESTEEM_ADDRESS")
	Port = os.Getenv("SELFESTEEM_PORT")

	DBUser = os.Getenv("SELFESTEEM_DB_USER")
	DBName = os.Getenv("SELFESTEEM_DB_NAME")
	DBHost = os.Getenv("SELFESTEEM_DB_HOST")
	DBPort = os.Getenv("SELFESTEEM_DB_PORT")
}
