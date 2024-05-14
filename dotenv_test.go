package godotenv

import (
	"fmt"
	"testing"
)

type Config struct {
	Debug bool `env:"DEBUG"`
	App   struct {
		Name string `env:"APP_NAME"`
		Port int16  `env:"APP_PORT"`
	}
	Database struct {
		Name string `env:"DB_NAME"`
		Port int16  `env:"DB_PORT"`
		User string `env:"DB_USER"`
		Pass string `env:"DB_PASS"`
	}
	IPs []string `env:"IPS"`
	IDs []int64  `env:"IDS"`
}

func TestLoad(t *testing.T) {
	var cfg Config
	err := Load(&cfg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("config %+v \n", cfg)
}
