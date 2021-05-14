package discord_messages

import (
	"fmt"
	"git-good-discord/utils"
	"io/ioutil"
	"reflect"
)

var languageFiles = make(map[string]commands)

// ReloadLanguageFiles reloads the language files
func ReloadLanguageFiles() error {
	var err error
	tempLanguageFiles, err := getLanguageFiles()
	if err != nil {
		return err
	}
	// Get the english language files, that will be used for substitution.
	englishFile := tempLanguageFiles["english"]
	english := reflect.ValueOf(&englishFile)
	// Loop through all languages (except english)
	for languageName := range tempLanguageFiles {
		if languageName != "english" {
			//Set language file to be the language being checked
			lang := tempLanguageFiles[languageName]
			// Get element for the commands struct
			commandsType := reflect.ValueOf(&lang).Elem()
			//Loop through each fields
			for i := 0; i < commandsType.NumField(); i++ {
				// Get the interface of each of the fields
				command := commandsType.Field(i).Interface()
				// Create a reflection of the interface
				e := reflect.ValueOf(command)
				// Loop through it if it is a struct
				if reflect.ValueOf(command).Kind() == reflect.Struct {
					for j := 0; j < e.NumField(); j++ {
						// If it is empty
						if e.Field(j).String() == "" {
							//Get name of struct, i.e ping
							cmdStruct := commandsType.Field(i).Type().Name()
							//Get entry, i.e roleNotFound
							cmdEntry := e.Type().Field(j).Name

							// Get the struct of the english language pack
							englishOuterStruct := english.Elem().FieldByName(cmdStruct).Interface()
							// Get the actual value from the english language pack
							englishValue := reflect.ValueOf(englishOuterStruct).FieldByName(cmdEntry)

							// Get the indirect reflection of lang
							tempLang := reflect.Indirect(reflect.ValueOf(&lang))
							// Set the value to be the english value
							tempLang.FieldByName(cmdStruct).FieldByName(cmdEntry).SetString(fmt.Sprintf("%v", englishValue))

							// Replace the language pack file with updated info
							tempLanguageFiles[languageName] = lang
						}
					}
				}
			}
		}
	}
	languageFiles = tempLanguageFiles

	return nil
}

// LoadLanguageFiles will load the language files and write to error channel if
// an error occurs
func LoadLanguageFiles(errorChan chan error) {
	err := ReloadLanguageFiles()
	if err != nil {
		errorChan <- err
	}
}

// getLanguageFiles will get language files for the given of commands
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

// parseLanguageFile will parse the given language file to commands
func parseLanguageFile(languageFileName string) (commands, error) {
	var language commands
	err := utils.FileToInterface(languageFileName, &language)
	return language, err
}

// getLanguage will get the commands for the given language
func getLanguage(language string) commands {
	if lang, ok := languageFiles[language]; ok {
		return lang
	}
	return languageFiles["english"]
}

// IsLanguage checks if language provided exists
func IsLanguage(language string) bool {
	if language != "english" {
		if _, ok := languageFiles[language]; !ok {
			return false
		}
	}
	return true
}
