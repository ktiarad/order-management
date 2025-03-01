package authentication

import (
	"context"
	"errors"
	"net/http"
	"order-management/internal/helper"
	"order-management/internal/model"
	"order-management/internal/repository"
)

type AuthenticationServices interface {
	SignIn(
		/*req*/ ctx context.Context, request model.User) (
		/*res*/ httpCode int, err error,
	)
	LogIn(
		/*req*/ ctx context.Context, request model.User) (
		/*res*/ token string, httpCode int, err error,
	)
}

type Service struct {
	Repo repository.Repository
}

func New(repo repository.Repository) AuthenticationServices {
	return &Service{
		Repo: repo,
	}
}

func (x *Service) SignIn(
	/*req*/ ctx context.Context, request model.User) (
	/*res*/ httpCode int, err error,
) {
	request.Password = helper.HashPass(request.Password)

	err = x.Repo.CreateUser(ctx, request)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (x *Service) LogIn(
	/*req*/ ctx context.Context, request model.User) (
	/*res*/ token string, httpCode int, err error,
) {
	user, err := x.Repo.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	if isAuthenticated := helper.ComparePass([]byte(user.Password), []byte(request.Password)); !isAuthenticated {
		return "", http.StatusUnauthorized, errors.New("wrong email or password")
	}

	token, err = helper.GenerateToken(request.Email, user.ID)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return token, http.StatusOK, nil
}
