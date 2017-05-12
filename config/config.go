package config

import (
	"os"
	"time"
	"log"
	"path/filepath"
	"bufio"
)

const (
	VerboseEnvVar = "API_VERBOSE" // set to "true" if verbose errors should be returned

	LogPoolSize = 10 // used to make a pool of items to send to dynamo
	LogChanSize = 100 // used to buffer the log chan size

	WorkerPoolSize = 3

	APPLICATION_PORT = "8080"

	JWT_EXPIRATION_HOURS = 13
)

func IsVerbose() bool {
	verbose := os.Getenv(VerboseEnvVar)
	if verbose == "true" {
		return true
	}
	return false
}

func JWTLifeTime() time.Duration {
	return time.Hour * JWT_EXPIRATION_HOURS
}

func GetSecretKey() (key []byte, err error) {
	homeDir := os.Getenv("HOME") // *nix
	if homeDir == "" {               // Windows
		homeDir = os.Getenv("USERPROFILE")
	}
	if homeDir == "" {
		log.Fatal("Can not find HOME env directory")
	}

	path := filepath.Join(homeDir, ".api-skeleton", "super-secret-key.dat")
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	key, err = buf.ReadBytes(0x00)
	return
}