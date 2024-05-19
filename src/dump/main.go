package dump

import (
	"backup_system/configs"
	"backup_system/log"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func CreateDump(database configs.Database) (string, error) {
	log.Log(fmt.Sprint("Criando dump da database:", database.Name))

	dateNow := time.Now().Format("2006-01-02:15-04-05")
	fileName := "/app/dumps/" + database.Name + "/" + database.Name + "-" + dateNow + ".sql"

	err1 := createDirIfNotExist(fileName)
	if err1 != nil {
		return "", err1
	}

	cmd := exec.Command("mysqldump", "-h"+database.Host, "-P "+database.Port, "-u"+database.User, "-p"+database.Password, database.Name, "--result-file="+fileName)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err2 := cmd.Run()
	if err2 != nil {
		return "", fmt.Errorf("%s: %s", err2, stderr.String())
	}

	log.Log(fmt.Sprint("O dump do banco de dados foi salvo no arquivo: ", fileName))
	return fileName, nil
}

func createDirIfNotExist(fileName string) error {
	dir := filepath.Dir(fileName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}
