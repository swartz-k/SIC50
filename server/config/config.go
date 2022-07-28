package config

import (
	"os"
	"path"
)

var Cfg *Config

type Config struct {
	Addr string `json:"addr"`

	MysqlHost     string `json:"mysql_host"`
	MysqlDatabase string `json:"mysql_database"`
	MysqlUser     string `json:"mysql_user"`
	MysqlPassword string `json:"mysql_password"`
	// local dev or single deploy
	SqlitePath string `json:"sqlite_path"`
	// upload image save path
	UploadPath string `json:"upload_path"`
}

// init from config.yaml or env
func init() {
	// fixme
	pwd, _ := os.Getwd()
	if Cfg == nil {
		sqliteDB := os.Getenv("SIC50_SQLITE_PATH")
		if sqliteDB == "" {
			sqliteDB = path.Join(pwd, "server.sqlite")
		}
		p := path.Join(pwd, "data")
		Cfg = &Config{
			Addr:       "127.0.0.1:9097",
			SqlitePath: sqliteDB,
			UploadPath: p,
		}
	}
}
