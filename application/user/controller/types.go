package controller

import (
	"github.com/google/uuid"
	"healthRoutine/application/domain/user"
)

type signUpData struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

func domainModelToSignUpResp(model *user.DomainModel) signUpData {
	return signUpData{
		Email:    model.Email,
		Nickname: model.Nickname,
	}
}

type signInData struct {
	Token    string `json:"token"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

func domainModelToSignInResp(model *user.DomainModel, token string) signInData {
	return signInData{
		Token:    token,
		Email:    model.Email,
		Nickname: model.Nickname,
	}
}

type profileData struct {
	Id           uuid.UUID `json:"id"`
	Nickname     string    `json:"nickname"`
	ProfileImage string    `json:"profileImage"`
}

func domainModelToProfileResp(model *user.DomainModel) profileData {
	return profileData{
		Id:           model.ID,
		Nickname:     model.Nickname,
		ProfileImage: model.ProfileImg,
	}
}
