package userRepository

import (
	"bt_auth/internal/model"
	"bt_auth/internal/repository"
	"bt_auth/internal/repository/userRepository/converter"
	"bt_auth/internal/repository/userRepository/model"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password_hash"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, user *model.UserInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(user.Name, user.Email, user.Password, user.Role).
		Suffix("RETURNING id")

	var userID int64
	sql, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	err = r.db.QueryRow(ctx, sql, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	sql, args, err := builder.ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var user modelRepo.User
	user.Info = &modelRepo.Info{}
	err = r.db.QueryRow(ctx, sql, args...).
		Scan(&user.ID, &user.Info.Name, &user.Info.Email, &user.Info.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Update(ctx context.Context, id int64, user *model.UserInfo) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Set(updatedAtColumn, sq.Expr("NOW()"))

	if user.Name != "" {
		builder = builder.Set(nameColumn, user.Name)
	}
	if user.Email != "" {
		builder = builder.Set(emailColumn, user.Email)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
