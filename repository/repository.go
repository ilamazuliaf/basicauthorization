package repository

import (
	"casbin/models"
	"context"
	"database/sql"
	"sync"
)

type (
	mysqlConfig struct {
		sync.Mutex
		DB *sql.DB
	}

	Repository interface {
		GetUser(ctx context.Context, username, password string) (*models.User, error)
		GetPersons(ctx context.Context) ([]*models.Person, error)
	}
)

const (
	getUser   = `select * from user where username = ? and password = ?`
	getPerson = `select nama, alamat, umur from person`
)

func NewRepositoryConfig(db *sql.DB) Repository {
	return &mysqlConfig{DB: db}
}

func (m *mysqlConfig) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	user := &models.User{}
	if err := m.DB.QueryRowContext(ctx, getUser, username, password).Scan(
		&user.Username, &user.Password, &user.Role,
	); err != nil {
		return nil, models.ErrUserNotFound
	}
	return user, nil
}

func (m *mysqlConfig) GetPersons(ctx context.Context) ([]*models.Person, error) {
	rows, err := m.DB.QueryContext(ctx, getPerson)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	person := make([]*models.Person, 0)
	for rows.Next() {
		P := &models.Person{}
		if err := rows.Scan(
			&P.Name, &P.Address, &P.Age,
		); err != nil {
			return nil, err
		}
		person = append(person, P)
	}
	return person, nil
}
