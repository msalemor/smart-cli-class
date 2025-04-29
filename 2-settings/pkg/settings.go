package pkg

import (
	"encoding/json"
	"os"
	"sync"
)

type Settings struct {
	Endpoint     string `json:"endpoint"`
	APIKey       string `json:"api_key"`
	Model        string `json:"model"`
	SystemPrompt string `json:"system_prompt"`
	Prompt       string `json:"prompt"`
}

var (
	settingsInstance *Settings
	once             sync.Once
)

func GetSettings() *Settings {
	once.Do(func() {
		settingsInstance = &Settings{}

		// TODO: load from where the executable is running
		data, err := os.ReadFile("/home/alex/temp/mytool/mytool.json")
		if err != nil {
			panic("Failed to read openai.json: " + err.Error())
		}

		if err := json.Unmarshal(data, settingsInstance); err != nil {
			panic("Failed to parse openai.json: " + err.Error())
		}

		if settingsInstance.Endpoint == "" || settingsInstance.APIKey == "" || settingsInstance.Model == "" || settingsInstance.SystemPrompt == "" {
			panic("Missing required environment variables ENDPOINT, API_KEY, MODEL or SYSTEM_PROMPT in openai.json")
		}
	})
	return settingsInstance
}
