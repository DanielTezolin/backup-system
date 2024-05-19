package main

import (
	"backup_system/configs"
	"backup_system/dump"
	"backup_system/log"
	"backup_system/notification"
	"backup_system/uploader"
	"fmt"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("0 3 * * *", initBackup)
	c.Start()

	select {}
}

func initBackup() {
	configs := configs.ReturnDatabases()
	erros := 0

	log.Log("---------------- Iniciando backup ----------------")

	for _, database := range configs.Databases {
		dumpPath, err := dump.CreateDump(database)
		if err != nil {
			erros++
			log.Log(fmt.Sprint(err))
			notification.SendTeamsNotification(fmt.Sprint("Erro ao fazer dump. ERRO:", err), "error", database.Name)
			continue
		}

		err = uploader.UploadFileToS3(dumpPath, database.Name)
		if err != nil {
			erros++
			log.Log(fmt.Sprint(err))
			notification.SendTeamsNotification(fmt.Sprint("Erro ao fazer upload na AWS. ERRO:", err), "error", database.Name)
			continue
		}
	}

	if erros > 0 {
		notification.SendTeamsNotification(fmt.Sprintf("Backup realizado com %d erros.", erros), "error", "")
		log.Log(fmt.Sprintf("---------------- Backup realizado com %d erros. ----------------", erros))
		return
	}

	log.Log("---------------- Backup realizado com sucesso. ----------------")
	notification.SendTeamsNotification("Backup realizado com sucesso.", "success", "")
}
