package godotenv

import (
	"fmt"
	"os"
)

// github.com/golobby/dotenv
import "github.com/golobby/dotenv"

func Load(config interface{}) error {
	file, err := os.Open(".env")
	if err != nil {
		return fmt.Errorf("open .env file error [%s]", err)
	}
	dec := dotenv.NewDecoder(file)
	err = dec.Decode(config)
	if err != nil {
		return fmt.Errorf("decode .env to structure error [%s]", err)
	}
	return nil
}
