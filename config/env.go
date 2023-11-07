package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dev-anderson/AnderBank/schemas"
)

func setEnv(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ports := strings.SplitN(line, "=", 2)
		if len(ports) == 2 {
			key := ports[0]
			value := ports[1]
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func LoadEnv() schemas.Env {
	err := setEnv(".env")
	if err != nil {
		fmt.Errorf("Error ao carregar o arquivo .env", err)
	}
	return schemas.Env{
		Host:     os.Getenv("anderbank-hostdb"),
		Port:     os.Getenv("anderbank-portdb"),
		User:     os.Getenv("anderbank-userdb"),
		Password: os.Getenv("anderbank-passworddb"),
		DBName:   os.Getenv("anderbank-dbname"),
		PortHttp: os.Getenv("anderbank-porthttp"),
	}
}
