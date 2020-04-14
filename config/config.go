package config

import (
	"log"
	"os"
)

// TODO: Herokuにこれらの環境変数を適切な値で設定する
var (
	isMockEnabled string
)

func init() {
	isMockEnabled = os.Getenv("ENABLE_MOCK")
}

// IsMockMode returns true when env ENABLE_MOCK is not empty.
func IsMockMode() bool {
	log.Println("isMockEnabled:", isMockEnabled)
	return isMockEnabled != ""
}
