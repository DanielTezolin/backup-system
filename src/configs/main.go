package configs

import (
	"backup_system/log"
	"encoding/json"
	"fmt"
	"os"
)

type Notification struct {
	Active       bool   `json:"active"`
	TeamsWebhook string `json:"teamsWebhook"`
}

type AWS struct {
	Active   bool   `json:"active"`
	S3Bucket string `json:"s3Bucket"`
	S3Region string `json:"s3Region"`
	S3Key    string `json:"s3Key"`
	S3Secret string `json:"s3Secret"`
}

type Database struct {
	Name     string `json:"dbname"`
	User     string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type Config struct {
	Databases    []Database   `json:"databases"`
	AWS          AWS          `json:"AWSupload"`
	Notification Notification `json:"notifications"`
}

func validateConfigDatabase(config Config) bool {
	for _, database := range config.Databases {
		if database.Name == "" || database.User == "" || database.Password == "" || database.Host == "" || database.Port == "" {
			return false
		}
	}
	return true
}

func validateAWSCredentials(aws AWS) bool {
	if !aws.Active {
		return true
	}

	if aws.S3Bucket == "" || aws.S3Region == "" || aws.S3Key == "" || aws.S3Secret == "" {
		return false
	}
	return true
}

func validateNotification(notification Notification) bool {
	if !notification.Active {
		return true
	}

	if notification.TeamsWebhook == "" {
		return false
	}

	return true
}

func ReturnDatabases() Config {
	configPath := "/app/config/config.json"
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Log(fmt.Sprint("Erro para ler o arquivo de configuracao", err))
		return Config{}
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Log(fmt.Sprint("Erro Unmarshal arquivo de configuracao", err))
		panic(err)
	}

	if !validateConfigDatabase(config) {
		log.Log("Erro ao validar banco de dados no arquivo de configuração.")
		panic(err)
	}
	return config
}

func ReturnAWSCredentials() AWS {
	configPath := "/app/config/config.json"
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Log(fmt.Sprint("Erro para ler o arquivo de configuracao", err))
		panic(err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Log(fmt.Sprint("Erro Unmarshal arquivo de configuracao", err))
		panic(err)
	}

	if !validateAWSCredentials(config.AWS) {
		log.Log("Erro ao validar credenciais da AWS no arquivo de configuração.")
		panic(err)
	}

	return config.AWS
}

func ReturnNotificationConfig() Notification {
	configPath := "/app/config/config.json"
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Log(fmt.Sprint("Erro para ler o arquivo de configuracao", err))
		panic(err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Log(fmt.Sprint("Erro Unmarshal arquivo de configuracao", err))
		panic(err)
	}

	if !validateNotification(config.Notification) {
		log.Log("Erro ao validar webhook do Teams no arquivo de configuração.")
		panic(err)
	}

	return config.Notification
}
