package repositories

import (
	"database/sql"
	"hexgo/internal/core/domain"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type userMysqlRepository struct {
	db *sql.DB
}

func NewUserMysqlRepository() *userMysqlRepository {
	connectionString := "user:passwd@/db_name"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		db.Close()
		log.Fatal(err)
	}

	return &userMysqlRepository{
		db: db,
	}
}

func (repo userMysqlRepository) Create(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (repo userMysqlRepository) List() ([]domain.User, error) {
	return []domain.User{}, nil
}

func (repo userMysqlRepository) Get(id string) (domain.User, error) {
	return domain.User{}, nil
}
