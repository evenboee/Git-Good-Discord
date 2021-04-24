package utils

import (
	"encoding/json"
	"google.golang.org/api/option"
	"io/ioutil"
	"os"
)

func GetFirestore() option.ClientOption {
	return option.WithCredentialsFile("./firebase.json")
}

func GetDiscordToken() (DiscordDetails, error) {
	var discordDetails DiscordDetails
	err := fileToInterface("discord.json", &discordDetails)
	return discordDetails, err
}

func GetFirestoreDetails() (FirestoreDetails, error) {
	var firestoreDetails FirestoreDetails
	err := fileToInterface("firestore.json", &firestoreDetails)
	return firestoreDetails, err
}

func fileToInterface(fileName string, data interface{}) error {
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
