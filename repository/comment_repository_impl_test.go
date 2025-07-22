package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_6_7__pzn_go_lang_database "go-lang-database"
	"go-lang-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(_6_7__pzn_go_lang_database.GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Email:   "test@mail.com",
		Comment: "Test comment",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}

func TestFindById(t *testing.T) {
	commentRepositoryImpl := NewCommentRepositoryImpl(_6_7__pzn_go_lang_database.GetConnection())
	ctx := context.Background()
	result, err := commentRepositoryImpl.FindById(ctx, 50)

	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(_6_7__pzn_go_lang_database.GetConnection())
	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
