package config

import (
	"github.com/BioChemML/SIC50/server/utils/log"
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
	// workdir  upload image save path
	WorkDir    string `json:"upload_path"`
	UploadDir  string `json:"upload_dir"`
	PythonPath string `json:"python_path"`
}

// init from config.yaml or env
func init() {
	workdir := os.Getenv("SIC50_WORKDIR")
	if workdir == "" {
		workdir, _ = os.Getwd()
	}
	if Cfg == nil {
		sqliteDB := os.Getenv("SIC50_SQLITE_PATH")
		if sqliteDB == "" {
			sqliteDB = path.Join(workdir, "server.sqlite")
		}
		Cfg = &Config{
			Addr:       "127.0.0.1:9097",
			SqlitePath: sqliteDB,
			WorkDir:    workdir,
			UploadDir:  path.Join(workdir, "data"),
		}
	}
	pythonPath := os.Getenv("SIC50_PYTHON_PATH")
	if pythonPath != "" {
		Cfg.PythonPath = pythonPath
	} else {
		Cfg.PythonPath = "python3.7"
	}
	log.Info("config %+v \n", Cfg)
}
