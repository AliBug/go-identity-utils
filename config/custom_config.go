package config

import (
	"log"

	"github.com/spf13/viper"
)

// ReadCustomStringConfig - return custom sect string value
func ReadCustomStringConfig(sect string) string {
	// 实际读取相应参数
	conf := viper.GetString(sect)

	if conf == "" {
		log.Fatalf("Field value of %s is required", sect)
	}

	return conf
}

// ReadCustomIntConfig - return custom sect int value
func ReadCustomIntConfig(sect string, allowZero bool) int {
	conf := viper.GetInt(sect)
	if allowZero {
		return conf
	}
	if conf == 0 {
		log.Fatalf("Field value of %s isn't allowed to be 0", sect)
	}
	return conf
}
