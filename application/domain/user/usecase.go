package user

import "context"

type SignUpParams struct {
	Nickname string
	Password string
	Email    string
}

type SignUpUseCase interface {
	Use(ctx context.Context, params SignUpParams) (*DomainModel, error)
}

type SignInUseCase interface {
	Use(ctx context.Context, email, password string) (*DomainModel, string, error)
}

type EmailValidationUseCase interface {
	Use(ctx context.Context, email string) error
}
