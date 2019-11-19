package env

import (
	"github.com/joho/godotenv"
	"os"
	"path"
	"runtime"
	"strconv"
)

type Env struct {
	DbHost       string
	DbPort       uint64
	DbName       string
	DbUser       string
	DbPass       string
	Port         uint64
	JwtSecretKey string
}

func GetEnv() Env {
	_, callerFileName, _, ok := runtime.Caller(0)

	if !ok {
		panic("No caller information")
	}

	dir := path.Dir(callerFileName)

	uri := path.Join(dir, "../.env")

	err := godotenv.Load(uri)

	if err != nil {
		panic(err)
	}

	dbPort, err := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 16)

	if err != nil {
		panic(err)
	}

	port, err := strconv.ParseUint(os.Getenv("PORT"), 10, 16)

	if err != nil {
		panic(err)
	}

	return Env{
		DbHost:       os.Getenv("DB_HOST"),
		DbPort:       dbPort,
		DbName:       os.Getenv("DB_NAME"),
		DbUser:       os.Getenv("DB_USER"),
		DbPass:       os.Getenv("DB_PASS"),
		Port:         port,
		JwtSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}
