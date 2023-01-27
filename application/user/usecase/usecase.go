package usecase

import (
	"healthRoutine/application/domain/user"
)

type UserUseCases struct {
	user.SignUpUseCase
	user.SignInUseCase
	user.EmailValidationUseCase
	user.GetProfileUseCase
	user.UploadTemporaryProfileUseCase
	user.UpdateProfileUseCase
	user.GetBadgeUseCase
	user.GetLatestBadgeUseCase
}
