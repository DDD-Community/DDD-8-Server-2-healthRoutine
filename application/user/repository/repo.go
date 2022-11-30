package repository

import (
	"context"
	"database/sql"
	"healthRoutine/application/domain/user"
	entity "healthRoutine/pkgs/database/sqlc/user"
)

func NewUserRepository(db *sql.DB) user.Repository {
	preparedQuery, err := entity.Prepare(context.Background(), db)
	if err != nil {
		panic(err)
	}
	return &repo{
		db:            db,
		query:         entity.New(db),
		preparedQuery: preparedQuery,
	}
}

type repo struct {
	db            *sql.DB
	query         *entity.Queries
	preparedQuery *entity.Queries
}
