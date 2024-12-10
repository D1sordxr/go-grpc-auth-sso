package handlers

import (
	"context"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/commands"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/application/persistence"
	"github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockDAO struct {
	mock.Mock
}

func (t *MockDAO) Register(ctx context.Context, tx interface{}, entity entity.User) error {
	args := t.Called(ctx, entity, tx)
	return args.Error(0)
}

func (t *MockDAO) Load(ctx context.Context, email string) (commands.User, error) {
	args := t.Called(ctx, email)
	return commands.User{}, args.Error(0)
}

func (t *MockDAO) Exists(ctx context.Context, email string) error {
	args := t.Called(ctx, email)
	return args.Error(0)
}

type TestUoW struct {
}

func (t *TestUoW) Begin(ctx context.Context) (interface{}, error) {
	return nil, nil
}
func (t *TestUoW) Commit(ctx context.Context) error {
	return nil
}
func (t *TestUoW) Rollback(ctx context.Context) error {
	return nil
}

type TestUoWManager struct {
}

func (t *TestUoWManager) GetUoW() persistence.UoW {
	return &TestUoW{}
}
