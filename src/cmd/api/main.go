package main

import (
	"os"

	"github.com/kenz68/cassandra-go/src/cmd/utils"
	api "github.com/kenz68/cassandra-go/src/pkg/api"
	log "github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func main() {

	config := read()

	url := os.Getenv("Cassandra_URL")
	if url != "" {
		config.Database.Url = url
	}

	log.Info("Cassandra_URL", url)

	api.Initilize(config)
}

// read configuration file from config.yml via Viper
func read() utils.Configuration {
	//Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var config utils.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Errorf("Unable to decode into struct, %v", err)
	}

	return config
}
