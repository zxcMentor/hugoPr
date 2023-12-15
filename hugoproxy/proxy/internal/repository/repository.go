package repository

import (
	"context"
	"database/sql"
	"proxy/internal/service/models"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) error
	GetByID(ctx context.Context, id string) (models.User, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, c models.Conditions) ([]models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user models.User) error {
	// тут реализация
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (models.User, error) {
	return models.User{}, nil
}

func (r *userRepository) Update(ctx context.Context, user models.User) error {
	// тут реализация
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	// тут реализация
	return nil
}

func (r *userRepository) List(ctx context.Context, c models.Conditions) ([]models.User, error) {
	// тут реализация
	return []models.User{}, nil
}
