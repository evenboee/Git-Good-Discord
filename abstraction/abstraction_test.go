package abstraction

import (
	"git-good-discord/database/database_mocks"
	"git-good-discord/discord/discord_mocks"
	"git-good-discord/gitlab/gitlab_mocks"
	"git-good-discord/gitlab/gitlab_structs"
	"github.com/golang/mock/gomock"
	"strconv"
	"testing"
)

func TestImplementation_HandleGitlabNotification(t *testing.T) {
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()

	// Initialize mocks
	discordMock := discord_mocks.NewMockInterface(mockctrl)
	databaseMock := database_mocks.NewMockDatabase(mockctrl)
	databaseConnectionMock := database_mocks.NewMockDatabaseConnection(mockctrl)
	gitlabMock := gitlab_mocks.NewMockInterface(mockctrl)

	// Initialize implementation
	implementation := Implementation{
		DiscordService:  discordMock,
		DatabaseService: databaseMock,
		GitlabService:   gitlabMock,
	}

	// Initialize required variables
	authorAndAssigneeUser := gitlab_structs.User{
		ID:       700,
		Name:     "nils",
		Username: "nils01",
		Email:    "nils@example.com",
	}

	webhookNotification := gitlab_structs.WebhookNotification{
		ObjectKind:       "issue",
		Project:          gitlab_structs.Project{
			URL:         "https://git.gvk.idi.ntnu.no/simen_bai/testing-git-good-discord-group-1",
			ID:          1965,
		},
		User:             gitlab_structs.User{
			ID:       900,
			Username: "roger",
		},
		ObjectAttributes: gitlab_structs.ObjectAttributes{
			AssigneeID:   authorAndAssigneeUser.ID,
			AuthorID:     authorAndAssigneeUser.ID,
		},
	}

	gitlabInstance := "git.gvk.idi.ntnu.no"
	discordChannelID := "831517371185102848"

	// Expected invocations
	databaseMock.EXPECT().GetConnection().MinTimes(1).Return(databaseConnectionMock)

	// Will lookup AssigneeID and AuthorID
	// But since they are the same, we are only expecting to invoke this ONCE
	gitlabMock.EXPECT().GetUserByID(webhookNotification.Project.URL, 700).Times(1).Return(authorAndAssigneeUser, nil)

	// Check that subscriptions are being fetched for each unique gitlab username
	databaseConnectionMock.EXPECT().GetSubscribers(discordChannelID, gitlabInstance, strconv.Itoa(webhookNotification.Project.ID), webhookNotification.User.Username).Times(1)
	databaseConnectionMock.EXPECT().GetSubscribers(discordChannelID, gitlabInstance, strconv.Itoa(webhookNotification.Project.ID), authorAndAssigneeUser.Username).Times(1)

	discordMock.EXPECT().SendMessage(gomock.Any()).Times(1)

	implementation.HandleGitlabNotification(webhookNotification, discordChannelID)
}