package usecase

import (
	"context"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"healthRoutine/application/domain/user"
	"healthRoutine/cmd"
	"healthRoutine/pkgs/log"
	"time"
)

var _ user.SignInUseCase = (*signInUseCaseImpl)(nil)

func SignInUseCase(repo user.Repository) user.SignInUseCase {
	return &signInUseCaseImpl{
		Repository: repo,
	}
}

type signInUseCaseImpl struct {
	user.Repository
}

func (u *signInUseCaseImpl) log() *zap.SugaredLogger {
	return log.Get().Named("SIGN_IN_USE_CASE")
}

func (u *signInUseCaseImpl) Use(ctx context.Context, email, password string) (resp *user.DomainModel, newToken string, err error) {
	logger := u.log()
	defer logger.Sync()

	config := cmd.LoadConfig()
	resp, err = u.Repository.GetByEmail(ctx, email)
	if err != nil {
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["user_id"] = resp.ID
	claims["exp"] = time.Now().Add(time.Hour * 8760).UnixMilli() // expired 임시 1년

	newToken, err = token.SignedString([]byte(config.SigningSecret))
	if err != nil {
		logger.Error(err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(password))
	if err != nil {
		logger.Error(err)
		return
	}

	return
}
