package repository

import (
	"database/sql"
	"log"

	"microgo/core/domain/user"
)

type UserRepository interface {
	Create(u user.User) error
	UserUpdate(u user.User) error
	FindByEmail(email string) (*user.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(u user.User) error {
	_, err := r.db.Exec(
		`INSERT INTO users (username, email, password, bio, profile_url, status)
         VALUES ($1, $2, $3, $4, $5, $6)`,
		u.Name, u.Email, u.Password, u.Bio, u.Profile_url, u.Status,
	)
	log.Fatal(err)

	return err

}

func (r *userRepository) FindByEmail(email string) (*user.User, error) {
	u := &user.User{}
	row := r.db.QueryRow(
		`SELECT id, username, email, password, bio, profile_url, status
         FROM users WHERE email=$1`, email,
	)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Bio, &u.Profile_url, &u.Status)
	log.Println(u)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userRepository) UserUpdate(u user.User) error {
	query := `
        UPDATE users 
        SET username=$1, email=$2, password=$3, bio=$4, profile_url=$5, status=$6
        WHERE id=$7
    `
	_, err := r.db.Exec(query, u.Name, u.Email, u.Password, u.Bio, u.Profile_url, u.Status, u.ID)
	return err
}
