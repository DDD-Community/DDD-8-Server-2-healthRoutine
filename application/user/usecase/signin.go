package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"healthRoutine/application/domain/user"
	"healthRoutine/cmd"
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

func (u *signInUseCaseImpl) Use(ctx context.Context, email, password string) (newToken string, err error) {
	config := cmd.LoadConfig()
	resp, err := u.Repository.GetByEmail(ctx, email)
	if err != nil || err == sql.ErrNoRows {
		panic(err)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["user_id"] = resp.ID
	claims["exp"] = time.Now().Add(time.Hour * 8760).UnixMilli() // expired 임시 1년

	newToken, err = token.SignedString([]byte(config.SigningSecret))
	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(password))
	if err != nil {
		fmt.Println("hash err")
		panic(err)
	}

	return
}
