package notification

import (
	"backup_system/configs"
	"backup_system/log"
	"bytes"
	"encoding/json"
	"net/http"
)

var config = configs.ReturnNotificationConfig()

func getColor(status string) string {
	if status == "success" {
		return "00e804"
	}
	return "ff0000"
}

func statusMenssageFactory(status string) string {
	if status == "success" {
		return "✅ Backup realizado com sucesso"
	}
	return "🔴 Erro ao realizar o backup"
}

func SendTeamsNotification(message string, status string, dbname string) error {
	if !config.Active {
		log.Log("Notificação desativada")
		return nil
	}

	notificationMessage := statusMenssageFactory(status)

	payload := map[string]interface{}{
		"@context":   "https://schema.org/extensions",
		"@type":      "MessageCard",
		"themeColor": getColor(status),
		"summary":    notificationMessage,
		"sections": []map[string]interface{}{
			{
				"activityImage": "https://eludico.com.br/favicon-96x96.png",
				"activityTitle": "Backup System",
			},
			{
				"activitySubtitle": notificationMessage,
			},
			{
				"activityTitle": "**Description**",
				"facts": []map[string]string{
					{
						"name":  "Database:",
						"value": dbname,
					},
					{
						"name":  "Mensagem:",
						"value": message,
					},
				},
			},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", config.TeamsWebhook, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
