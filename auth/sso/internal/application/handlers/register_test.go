package handlers

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSuccessRegisterUserHandle(t *testing.T) {
	fixedUserID := "123e4567-e89b-12d3-a456-426614174000"
	registerUserCommand := commands.RegisterUserCommand{}
	if err := faker.FakeData(&registerUserCommand); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	mockStorage := new(MockStorage)
	mockStorage.On("Exists", mock.Anything, mock.Anything).Return(nil)
	mockStorage.On("Register", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	registerUserCommand.Email = "testing@email.now"
	registerUserCommand.Password = "The_Hardest_Password"

	registerUser := NewRegisterUserHandler(mockStorage, &TestUoWManager{})
	userDTO, err := registerUser.Handle(ctx, registerUserCommand)
	if err != nil {
		t.Fatal(err)
	}

	if len(userDTO.UserID) != len(fixedUserID) {
		t.Fatalf("unexpected userID's length: %v, expected: %v", len(userDTO.UserID), len(fixedUserID))
	}
	t.Logf("userID's length %v is equal to expected %v", len(userDTO.UserID), len(fixedUserID))

	if userDTO.Email != registerUserCommand.Email {
		t.Fatalf("expected email: %s, got: %s", registerUserCommand.Email, userDTO.Email)
	}
	t.Logf("email: %s, got: %s", registerUserCommand.Email, userDTO.Email)

	mockStorage.AssertExpectations(t)
}
