package database

import (
	"fmt"
	"literature/models"
	"literature/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Literatur{},
		&models.Collection{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Filed")
	}

	fmt.Println("Migration Success")
}
