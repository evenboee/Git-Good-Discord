package utils

import (
	"encoding/json"
	"google.golang.org/api/option"
	"io/ioutil"
	"os"
)

// GetFirestore get Firestore credentials
func GetFirestore() option.ClientOption {
	return option.WithCredentialsFile("./firebase.json")
}

// GetDiscordToken gets the Discord Details
func GetDiscordToken() (DiscordDetails, error) {
	var discordDetails DiscordDetails
	err := FileToInterface("discord.json", &discordDetails)
	return discordDetails, err
}

// GetFirestoreDetails gets the Firestore Details
func GetFirestoreDetails() (FirestoreDetails, error) {
	var firestoreDetails FirestoreDetails
	err := FileToInterface("firestore.json", &firestoreDetails)
	return firestoreDetails, err
}

// FileToInterface reads file and unmarshalls it to interface
func FileToInterface(fileName string, data interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return err
	}

	return nil
}
