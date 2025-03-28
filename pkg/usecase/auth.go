package usecase

import (
	"context"

	"github.com/joejosephvarghese/message/server/pkg/domain"
	interfaces "github.com/joejosephvarghese/message/server/pkg/usecase/interface"
)

type authUseCase struct{}

func NewAuthUseCase() interfaces.AuthUseCase {
	return &authUseCase{}
}
func (c *authUseCase) UserSignUp(ctx context.Context, signUpDetails domain.User) error {
	return nil
}
