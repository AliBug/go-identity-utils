package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// ReadCasbinFilePath -
func ReadCasbinFilePath(sect string) string {
	// 实际读取相应参数
	filename := viper.GetString(fmt.Sprintf("%s.modelfile", sect))

	if filename == "" {
		log.Fatalf("Model file of %s is required", filename)
	}

	return filename
}
