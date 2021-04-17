package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const mongoURLTemplate = "mongodb://%s:%s@%s/%s"

// ReadMongoConfig - return mongodb conn url
func ReadMongoConfig(sect string) string {
	// 实际读取相应参数
	mongoUser := viper.GetString(fmt.Sprintf("%s.user", sect))
	mongoPass := viper.GetString(fmt.Sprintf("%s.pass", sect))
	mongoHost := viper.GetString(fmt.Sprintf("%s.host", sect))
	mongoDbName := viper.GetString(fmt.Sprintf("%s.database", sect))

	if mongoHost == "" || mongoDbName == "" {
		log.Fatalf("Mongo host or dbname of %s is required", sect)
	}

	return fmt.Sprintf(mongoURLTemplate, mongoUser, mongoPass, mongoHost, mongoDbName)
}
