package discord_messages

import (
	"git-good-discord/utils"
	"io/ioutil"
)

var languageFiles = make(map[string]commands)

func ReloadLanguageFiles() error {
	var err error
	tempLanguageFiles, err := getLanguageFiles()
	if err != nil {
		return err
	}
	languageFiles = tempLanguageFiles

	return nil
}

func LoadLanguageFiles(errorChan chan error) {
	err := ReloadLanguageFiles()
	if err != nil {
		errorChan <- err
	}
	currentLanguagePack = languageFiles["english"]
}

func getLanguageFiles() (map[string]commands, error) {
	files, err := ioutil.ReadDir("languages")
	if err != nil {
		return make(map[string]commands), err
	}
	commandsMap := make(map[string]commands)
	for _, file := range files {
		command, err := parseLanguageFile("languages/" + file.Name())
		if err == nil {
			commandsMap[command.Language] = command
		}
	}
	return commandsMap, nil
}

func parseLanguageFile(languageFileName string) (commands, error) {
	var language commands
	err := utils.FileToInterface(languageFileName, &language)
	return language, err
}
