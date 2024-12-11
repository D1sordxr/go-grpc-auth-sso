package handlers

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/vo"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/token"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/token/config"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

// TODO: login changes
func TestSuccessLoginUserHandle(t *testing.T) {
	loginUserCommand := commands.LoginUserCommand{}
	if err := faker.FakeData(&loginUserCommand); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	mockStorage := new(MockStorage)

	password, _ := vo.NewPassword("hashed_password")

	mockStorage.On("Load", mock.Anything, mock.Anything).Return(&commands.User{
		ID:       1,
		UserID:   vo.NewUserID().UserID,
		Email:    "testing@email.now",
		Password: password.GetHashedPassword(),
		AppID:    1,
	}, nil)

	cfg := config.TokenConfig{
		Key:      "secretKey",
		TokenTTL: time.Duration(420000),
	}
	tokenService := token.NewTokenService(&cfg)

	loginUserCommand.Email = "testing@email.now"
	loginUserCommand.Password = "hashed_password"
	loginUserCommand.AppID = 1

	loginUser := NewLoginUserHandler(mockStorage, &TestUoWManager{}, tokenService)
	userDTO, err := loginUser.Handle(ctx, loginUserCommand)
	if err != nil {
		t.Fatal(err)
	}

	if userDTO.Email != loginUserCommand.Email {
		t.Fatalf("expected email: %s, got: %s", loginUserCommand.Email, userDTO.Email)
	}
	t.Logf("email: %s, got: %s", loginUserCommand.Email, userDTO.Email)

	if userDTO.Token == "" {
		t.Fatal("expected token, got empty")
	}
	t.Logf("token: %s", userDTO.Token)

	mockStorage.AssertExpectations(t)
}
