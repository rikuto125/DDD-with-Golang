package config

import (
	"example.com/m/domain/model"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db  *gorm.DB
	err error
)

// DB接続
func Connect() *gorm.DB {
	// 実行環境取得
	env := os.Getenv("ENV")

	if "production" == env {
		env = "production"
	} else {
		env = "development"
	}

	// 環境変数取得
	godotenv.Load(".env." + env)
	godotenv.Load()

	// DB接続
	db, err = gorm.Open("mysql", os.Getenv("CONNECT"))

	//table作成
	db.AutoMigrate(&model.User{})

	//最初の時だけmysqlにログインして作るのがめんどくさかったら以下のコメントアウトを外してください
	//db.Create(&model.User{
	//	Id:   1,
	//	Name: "田中",
	//})

	if err != nil {
		fmt.Println(err)
	}
	//defer Close()
	return db
}

// DB終了
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
