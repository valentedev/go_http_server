package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Basic struct {
	ID        string `json:"id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type BasicModel struct {
	DB *sql.DB
}

func (bm BasicModel) Get(id string) (*Basic, error) {
	query := `
	SELECT * FROM basic;
	`

	var basic Basic

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := bm.DB.QueryRowContext(ctx, query, id).Scan(
		&basic.ID,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &basic, nil
}
