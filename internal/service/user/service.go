package user

import (
	"errors"

	"github.com/rafli024/mytodo-app/internal/contract"
	"github.com/rafli024/mytodo-app/internal/entities"
	"github.com/rafli024/mytodo-app/pkg/autils"
)

type Service struct {
	app  *contract.App
	repo contract.UserRepository
}

func InitUserService(a *contract.App) (svc contract.UserService) {
	r := NewUserRepository(a)

	svc = &Service{
		app:  a,
		repo: r,
	}

	return
}

// Register implements contract.UserService.
func (s *Service) Register(user entities.User) (err error) {
	hashedPassword, err := autils.HashAndSalt(user.Password)
	if err != nil {
		s.app.Logger.Error().Stack().Err(err).Msg("Failed to hash password")
		return err
	}

	user.Password = hashedPassword

	err = s.repo.Insert(user)

	return
}

// GetById implements contract.UserService.
func (s *Service) GetById(id string) (user entities.User, err error) {
	user, err = s.repo.FindById(id)

	return
}

// Login implements contract.UserService.
func (s *Service) Login(username string, password string) (user entities.User, err error) {
	user, err = s.repo.FindByUsername(username)
	if err != nil {
		// This will handle both "not found" and other database errors.
		return entities.User{}, errors.New("invalid username or password")
	}

	passwordsMatch, err := autils.ComparePasswords(user.Password, password)
	if err != nil {
		s.app.Logger.Error().Stack().Err(err).Msg("Failed during password comparison")
		return entities.User{}, errors.New("an error occurred during login")
	}

	if !passwordsMatch {
		return entities.User{}, errors.New("invalid username or password")
	}

	return user, nil
}

// GetByUsername implements contract.UserService.
func (s *Service) GetByUsername(username string) (user entities.User, err error) {
	user, err = s.repo.FindByUsername(username)

	return
}
