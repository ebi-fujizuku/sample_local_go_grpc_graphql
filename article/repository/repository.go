package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ebi-fujizuku/sample_local_go_grpc_graphql/article/common"
	myConf "github.com/ebi-fujizuku/sample_local_go_grpc_graphql/article/config"
	"github.com/ebi-fujizuku/sample_local_go_grpc_graphql/article/pb"
	_ "github.com/mattn/go-sqlite3"
)

type Repository interface{
	InsertArticle(context.Context,*pb.ArticleInput)(int64,error)
	SelectArticleByID(context.Context,int64)(*pb.Article,error)
	UpdateArticle(context.Context,int64,*pb.ArticleInput)error
	DeleteArticle(context.Context,int64)error
	SelectAllArticles()(*sql.Rows,error)
}

type sqliteRepo struct{
	db *sql.DB
}

func NewSqliteRepo()(Repository,error){
	common.PrintStart("")
	fmt.Println("repository.NewsqliteRepo")
	fmt.Println("  conf",myConf.Conf)
	fmt.Println("  Sqlite3_path",myConf.Conf.Sqlite3_path)
	if myConf.Conf.Sqlite3_path ==""{
		panic(errors.New("Sqlite3_path が 空文字です。/article/.env SQLITE3_PATHを確認してください。"))
	}
	db,err := sql.Open("sqlite3",myConf.Conf.Sqlite3_path)
	if err != nil{
		fmt.Println("  Failed DB_Open")
		return nil,err
	}
	fmt.Println("  DB_Open")

	//articlesテーブルを作成
	cmd := `CREATE TABLE IF NOT EXISTS articles(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		author  STRING,
		title   STRING,
		content STRING
		)`
		_,err = db.Exec(cmd)
		if err != nil{
			fmt.Println("  Failed DB_Create")
			return nil,err
		}
		fmt.Println("  DB_Create")

		common.PrintEnd("")
	return &sqliteRepo{db},nil
}

func (r *sqliteRepo) InsertArticle(ctx context.Context,input *pb.ArticleInput) (int64, error) {
	common.PrintStart("")
	fmt.Println("repository.InsertArticle")
	// Inputの内容(Author, Title, Content)をarticlesテーブルにINSERT
	cmd := "INSERT INTO articles(author, title, content) VALUES (?, ?, ?)"
	result, err := r.db.Exec(cmd, input.Author, input.Title, input.Content)
	if err != nil {
		return 0, err
	}
	fmt.Println("  Pass:insert")

	// INSERTした記事のIDを取得
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// INSERTした記事のIDを返す
	common.PrintEnd("")
	return id, nil
}

func (r *sqliteRepo)SelectArticleByID(ctx context.Context,id int64)(*pb.Article,error){
	// 該当IDの記事をSELECT
	cmd := "SELECT * FROM articles WHERE id = ?"
	row := r.db.QueryRow(cmd,id)
	var a pb.Article

	// SELECTした記事の内容を読み取る
	err := row.Scan(&a.Id, &a.Author, &a.Title, &a.Content)
	if err != nil{
		return nil,err
	}

	return &pb.Article{
		Id: a.Id,
		Author: a.Author,
		Title: a.Title,
		Content: a.Content,
	},nil
}

func (r *sqliteRepo)UpdateArticle(ctx context.Context,id int64,input *pb.ArticleInput)(error){
	// 該当IDをUPDATE
	cmd := "UPDATE articles SET author = ?, title = ?, content = ? WHERE id = ?"
	_, err := r.db.Exec(cmd, input.Author, input.Title, input.Content, id)
	if err != nil{
		return err
	}

	return nil
}

func (r *sqliteRepo)DeleteArticle(ctx context.Context,id int64)(error){
	// 該当IDをDELETE
	cmd := "DELETE FROM articles WHERE id = ?"
	_, err := r.db.Exec(cmd, id)
	if err != nil{
		return err
	}

	return nil
}
func (r *sqliteRepo)SelectAllArticles()(*sql.Rows,error){
	// 該当IDの記事をSELECT
	cmd := "SELECT * FROM articles"
	rows,err := r.db.Query(cmd)
	if err != nil{
		return nil,err
	}

	return rows,nil
}
