package repository_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"github.com/snowball-devs/backend-utec-inscriptions/repository"
	"github.com/snowball-devs/backend-utec-inscriptions/repository/mocks"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var repo *mocks.UserRepository

// Init mock instance to unit testing
func TestMain(m *testing.M) {
	repo = &mocks.UserRepository{}
	repository.SetUserRepository(repo)
	code := m.Run()
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {

	testCases := []struct {
		Name          string
		Input         models.User
		ExpectedError error
	}{
		{
			Name: "Success Create new user",
			Input: models.User{
				ID:          primitive.NewObjectID(),
				Email:       "estefany.lue99@gmail.com",
				Username:    "estefany.lue99",
				Password:    "vanillagolang123",
				Permissions: "manager",
				Disable:     false,
				CreatedAt:   time.Now(),
			},
			ExpectedError: nil,
		},
	}

	ctx := context.Background()
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.On("CreateUser", ctx, &tc.Input).Return("New user created", nil)
			_, err := repository.CreateUser(ctx, &tc.Input)
			if err != nil {
				t.Errorf("Create user incorrect, go %v want %v", tc.ExpectedError, err)
			}
		})
	}
}