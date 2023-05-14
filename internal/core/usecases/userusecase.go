package usecases

import (
	"hexgo/internal/core/domain"
	"hexgo/internal/core/ports"
)

type userUseCase struct {
	repo ports.UserRepository
}

func NewUserUseCase(repo ports.UserRepository) *userUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (u userUseCase) Create(user domain.User) (domain.User, error) {
	return u.repo.Create(user)
}

func (u userUseCase) List() ([]domain.User, error) {
	return u.repo.List()
}

func (u userUseCase) Get(id string) (domain.User, error) {
	return u.repo.Get(id)
}
