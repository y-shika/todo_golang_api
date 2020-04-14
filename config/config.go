package config

import (
	"os"
)

var (
	usage         string
	isMockEnabled string
	connectedDB   string
)

func init() {
	usage = os.Getenv("USAGE")
	isMockEnabled = os.Getenv("ENABLE_MOCK")
	connectedDB = os.Getenv("CONNECTED_DATABASE")
}

// GetUsage returns a current usage level.
func GetUsage() string {
	return usage
}

// IsMockMode returns true when env ENABLE_MOCK is not empty.
func IsMockMode() bool {
	return isMockEnabled != ""
}

// IsConnectedHerokuDB returns true when env CONNECTED_DATABASE is "heroku".
func IsConnectedHerokuDB() bool {
	return connectedDB == "heroku"
}

// IsLocal returns true when usage is "local".
func IsLocal() bool {
	return usage == "local"
}

// IsHeroku returns true when usage is "heroku".
// TODO: 本来ならdev, prodとかだと思うが、今回はそのような分け方してないので一旦Herokuとする
//       のちに良い名前に変えたい
func IsHeroku() bool {
	return usage == "heroku"
}
