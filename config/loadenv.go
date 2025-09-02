package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func loadENV() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
