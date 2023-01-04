package usecase

import (
	"healthRoutine/application/domain/user"
)

const (
	named = "USER_USE_CASE"
)

type UseCases struct {
	user.SignUpUseCase
	user.SignInUseCase
	user.EmailValidationUseCase
	user.GetProfileUseCase
	user.UploadTemporaryProfileUseCase
	user.UpdateProfileUseCase
}
