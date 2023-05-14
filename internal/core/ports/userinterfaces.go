package ports

import "hexgo/internal/core/domain"

type UserRepository interface {
	Create(user domain.User) (domain.User, error)
	List() ([]domain.User, error)
	Get(id string) (domain.User, error)
}

type UserUseCase interface {
	Create(user domain.User) (domain.User, error)
	List() ([]domain.User, error)
	Get(id string) (domain.User, error)
}
