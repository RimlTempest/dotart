package main

import (
	"fmt"

	// 独自パッケージ
	"api/constants"
	"api/domain/model"
	"api/router"

	// dbパッケージ
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// DBマイグレーション
	db := constants.Init()
	DBMigrate(db)

	router := router.SetRoutes(db)

	// APIサーバ起動
	router.Run(constants.API_SERVER)
}

// DBMigrate DBを生成
func DBMigrate(db *gorm.DB) *gorm.DB {
	if db.HasTable(&model.User{}) || db.HasTable(&model.Project{}) {
		fmt.Println("Migrateされたテーブルが存在します。")
	} else {
		fmt.Println("Migrateされたテーブルが存在しません。")
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{}, &model.Project{})
	fmt.Println("AUTO_MIGRATE : OK")
	return db
}
