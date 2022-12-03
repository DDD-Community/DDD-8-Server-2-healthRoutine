package usecase

import (
	"healthRoutine/application/domain/user"
)

type UseCases struct {
	user.SignUpUseCase
	user.SignInUseCase
}
