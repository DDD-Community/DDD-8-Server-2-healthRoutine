package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"healthRoutine/application/domain/user"
	entity "healthRoutine/pkgs/database/sqlc/user"
	"time"
)

var _ user.Repository = (*repo)(nil)

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

func (r *repo) Create(ctx context.Context, nickname, email, password string) (err error) {
	now := time.Now().UnixMilli()
	defaultImgUrl := "https://picsum.photos/536/354" // need change
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return
	}

	return r.preparedQuery.Create(ctx, entity.CreateParams{
		ID:         uuid.New(),
		Nickname:   nickname, // TODO: nickname 으로 변경
		Email:      email,
		Password:   string(hashPassword),
		ProfileImg: defaultImgUrl,
		CreatedAt:  now,
		UpdatedAt:  now,
	})
}

func (r *repo) GetByEmail(ctx context.Context, email string) (*user.DomainModel, error) {
	resp, err := r.preparedQuery.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &user.DomainModel{
		User: resp,
	}, err
}
