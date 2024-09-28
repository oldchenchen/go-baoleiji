package Config

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename,omitempty"`
	Maxsize    int    `json:"maxsize,omitempty"`
	Maxage     int    `json:"max_age,omitempty"`
	MaxBackups int    `json:"max_backups,omitempty"`
}
type DatabaseConfig struct {
	Driver          string `json:"driver"`
	Host            string `json:"host"`
	Port            string `json:"port"`
	Database        string `json:"database"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Charset         string `json:"charset"`
	MaximumConn     int    `json:"maximum_connection"`
	MaximumFreeConn int    `json:"maximum_free_connection"`
	TimeOut         int    `json:"timeout"`
}
type Configs struct {
	Mode                string `json:"mode" `
	Host                string `json:"host"`
	Port                int    `json:"port"`
	SecretKey           string `json:"secret_key"`
	*LogConfig          `json:"log"`
	*DatabaseConfig     `json:"database"`
	*jwt.StandardClaims `json:"jwt"`
}

var Conf = new(Configs)

func Initconfig(filepath string) error {
	//配置文件路径
	b, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, Conf)
}
