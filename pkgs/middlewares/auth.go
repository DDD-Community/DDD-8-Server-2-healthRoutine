package middlewares

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"healthRoutine/cmd"
	"healthRoutine/pkgs/errors/response"
	"strings"
)

func getToken(ctx *fiber.Ctx) (tokenString string, err error) {
	reqToken := ctx.Get("Authorization")
	if !strings.HasPrefix(reqToken, "Bearer") {
		err = response.ErrUnauthorized
		return "", response.ErrorResponse(ctx, err, nil)
	}

	tokenString = strings.TrimPrefix(reqToken, "Bearer ")

	return
}

func parseToken(tokenString string) (*jwt.Token, error) {
	cfg := cmd.LoadConfig()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("wrong signing method")
		}
		return []byte(cfg.SigningSecret), nil
	})
	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func AuthRequired() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		tokenString, err := getToken(ctx)
		if err != nil {
			err = response.ErrUnauthorized
			return response.ErrorResponse(ctx, err, nil)
		}

		token, err := parseToken(tokenString)
		if err != nil {
			err = response.ErrUnauthorized
			return response.ErrorResponse(ctx, err, nil)
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			err = response.ErrUnauthorized
			return response.ErrorResponse(ctx, err, nil)
		}

		return ctx.Next()
	}
}

func ExtractUserId(ctx *fiber.Ctx) (uuid.UUID, error) {
	tokenString, err := getToken(ctx)
	if err != nil {
		err = response.ErrUnauthorized
		return uuid.Nil, response.ErrorResponse(ctx, err, nil)
	}

	token, err := parseToken(tokenString)
	if err != nil {
		err = response.ErrUnauthorized
		return uuid.Nil, response.ErrorResponse(ctx, err, nil)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = response.ErrUnauthorized
		return uuid.Nil, response.ErrorResponse(ctx, err, nil)
	}

	userId, _ := uuid.Parse(claims["user_id"].(string))

	return userId, err

}
