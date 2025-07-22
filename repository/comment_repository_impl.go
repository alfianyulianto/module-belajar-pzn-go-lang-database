package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-lang-database/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepositoryImpl(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlExec := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := repo.DB.ExecContext(ctx, sqlExec, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	sqlExec := "SELECT id, email, comment FROM comments WHERE id=?"
	rows, err := repo.DB.QueryContext(ctx, sqlExec, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return comment, err
		}

		return comment, nil

	} else {
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sqlExec := "SELECT id, email, comment FROM comments"
	rows, err := repo.DB.QueryContext(ctx, sqlExec)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []entity.Comment{}
	for rows.Next() {
		comment := entity.Comment{}
		err = rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
