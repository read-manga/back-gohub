package repository

import (
	"database/sql"
	"log"

	"microgo/core/domain/repo"
	"microgo/core/domain/user"

	"github.com/lib/pq"
)

type UserRepository interface {
	Create(u user.User) error
	Update(u user.User) error
	UpdateRepositoryByUser(r repo.Repo) error
	FindByEmail(email string) (*user.User, error)
	FindById(id string) (*user.User, error)
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
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}





func (r *userRepository) UpdateRepositoryByUser(rRepo repo.Repo) error {
    existing := repo.Repo{}
    err := r.db.QueryRow(`
        SELECT name, about, tag, public, storage_s3, colaborators, pin 
        FROM repos 
        WHERE id=$1 AND user_id=$2
    `, rRepo.ID, rRepo.UserId).Scan(
        &existing.Name,
        &existing.About,
        pq.Array(&existing.Tag),
        &existing.Public,
        &existing.StorageS3,
        pq.Array(&existing.Colaborators),
        &existing.Pin,
    )
    if err != nil {
        return err
    }


    if rRepo.Name != "" {
        existing.Name = rRepo.Name
    }
    if rRepo.About != "" {
        existing.About = rRepo.About
    }
    if rRepo.Tag != nil {
        existing.Tag = rRepo.Tag
    }
    if rRepo.StorageS3 != "" {
        existing.StorageS3 = rRepo.StorageS3
    }
    if rRepo.Colaborators != nil {
        existing.Colaborators = rRepo.Colaborators
    }
    existing.Public = rRepo.Public
    existing.Pin = rRepo.Pin


    _, err = r.db.Exec(`
        UPDATE repos SET name=$1, about=$2, tag=$3, public=$4, storage_s3=$5, colaborators=$6, pin=$7
        WHERE id=$8 AND user_id=$9
    `, existing.Name, existing.About, pq.Array(existing.Tag), existing.Public, existing.StorageS3, pq.Array(existing.Colaborators), existing.Pin, rRepo.ID, rRepo.UserId)

    return err
}



func (r *userRepository) Update(u user.User) error {
	query := `
        UPDATE users 
        SET username=$1, email=$2, password=$3, bio=$4, profile_url=$5, status=$6
        WHERE id=$7
    `

	_, err := r.db.Exec(query, u.Name, u.Email, u.Password, u.Bio, u.Profile_url, u.Status, u.ID)
	return err
}


func (r *userRepository) FindById(id string) (*user.User, error) {
	u := &user.User{}


	row := r.db.QueryRow(`
		SELECT id, username, email, password, bio, profile_url, status
		FROM users
		WHERE id = $1
	`, id)

	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Bio,
		&u.Profile_url,
		&u.Status,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}


	rows, err := r.db.Query(`
		SELECT id, name, about, tag, public, storage_s3, user_id, colaborators, pin
		FROM repos
		WHERE user_id = $1
	`, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var repos []repo.Repo

	for rows.Next() {
		var rp repo.Repo

		err := rows.Scan(
			&rp.ID,
			&rp.Name,
			&rp.About,

			(*pq.StringArray)(&rp.Tag),
			&rp.Public,
			&rp.StorageS3,
			&rp.UserId,
			(*pq.StringArray)(&rp.Colaborators),
			&rp.Pin,
		)

		if err != nil {
			return nil, err
		}

		repos = append(repos, rp)
	}

	u.Repos = repos
	return u, nil
}
