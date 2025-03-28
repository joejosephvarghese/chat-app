package interfaces

import (
	"context"

	"github.com/joejosephvarghese/message/server/pkg/domain"
)

type AuthUseCase interface {
	//user
	UserSignUp(ctx context.Context, signUpDetails domain.User) error
}
