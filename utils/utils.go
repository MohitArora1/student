package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
)

type config struct {
	DatabaseHost string `mapstructure:"database_host"`
	DatabaseName string `mapstructure:"database_name"`
	DatabasePort int    `mapstructure:"database_port"`
}

var (
	//Config overall config settings
	Config config
)

//InitConfig setup configuration
func InitConfig() {

	log.SetFlags(log.LstdFlags | log.Llongfile)
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.ReadInConfig()
	viper.Unmarshal(&Config)

	log.Printf("\n\nCONFIGURATION\n")
	log.Printf("\n================================================================================\n")
	displayConfig := Config
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(displayConfig)
	log.Printf("\n================================================================================\n")
}

//GetDatabaseSession return session
func GetDatabaseSession() (session *mgo.Session, err error) {

	Host := []string{
		Config.DatabaseHost + ":" + strconv.Itoa(Config.DatabasePort),
	}
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: Host,
	})
	return
}

//WriteJSON return json result
func WriteJSON(w http.ResponseWriter, element interface{}) {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	w.Header().Set("Content-Type", "application/json")
	enc.Encode(element)
}
