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
	defaultImgUrl := "https://user-images.githubusercontent.com/2377807/209339362-fc391ce0-d7ab-4bc4-abaa-ae836ce031e7.png"
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return
	}

	return r.preparedQuery.Create(ctx, entity.CreateParams{
		ID:         uuid.New(),
		Nickname:   nickname,
		Email:      email,
		Password:   string(hashPassword),
		ProfileImg: defaultImgUrl,
		CreatedAt:  now,
		UpdatedAt:  now,
	})
}
func (r *repo) GetById(ctx context.Context, id uuid.UUID) (*user.DomainModel, error) {
	resp, err := r.preparedQuery.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user.DomainModel{
		User: resp,
	}, err
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

func (r *repo) CheckExistsByEmail(ctx context.Context, email string) (bool, error) {
	return r.preparedQuery.CheckExistsByEmail(ctx, email)
}

func (r *repo) CheckExistsByNickname(ctx context.Context, nickname string) (bool, error) {
	return r.preparedQuery.CheckExistsByNickname(ctx, nickname)
}

func (r *repo) UpdateProfileImgById(ctx context.Context, id uuid.UUID, url string) error {
	return r.preparedQuery.UpdateProfileImgById(ctx, entity.UpdateProfileImgByIdParams{
		ProfileImg: url,
		UpdatedAt:  time.Now().UnixMilli(),
		ID:         id,
	})
}

func (r *repo) UpdateNicknameById(ctx context.Context, id uuid.UUID, nickname string) error {
	return r.preparedQuery.UpdateNicknameById(ctx, entity.UpdateNicknameByIdParams{
		Nickname:  nickname,
		UpdatedAt: time.Now().UnixMilli(),
		ID:        id,
	})
}
