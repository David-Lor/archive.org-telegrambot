package settings

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	settingsFileEnvKey  = "SETTINGS_FILE"
	settingsFileDefault = "settings.yaml"
)

type Settings struct {
	AMQP       AMQPSettings       `yaml:"amqp"`
	Telegram   TelegramSettings   `yaml:"telegram"`
	Archiveorg ArchiveorgSettings `yaml:"archiveorg"`
}

func (settings *Settings) ValidateSettings() (err error) {
	settings.Archiveorg.ParsedTimeout, err = time.ParseDuration(settings.Archiveorg.RawTimeout)
	return
}

func LoadSettings(filePath string) (settings *Settings, err error) {
	if filePath == "" {
		filePath = getFilePathFromEnvOrDefault()
	}

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	var _settings Settings
	err = yaml.Unmarshal(fileData, &_settings)
	if err != nil {
		return
	}

	settings = &_settings
	err = settings.ValidateSettings()
	return
}

func getFilePathFromEnvOrDefault() string {
	path := os.Getenv(settingsFileEnvKey)
	if path == "" {
		path = settingsFileDefault
	}
	return path
}
