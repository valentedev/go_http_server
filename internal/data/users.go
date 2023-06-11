package data

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UserModel struct {
	DB *sql.DB
}

func (um UserModel) Get() ([]*User, error) {
	query := `
	SELECT * FROM users;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := um.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.ID,
			&user.Name,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
