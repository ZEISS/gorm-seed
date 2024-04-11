package main

import (
	"context"

	seed "github.com/zeiss/gorm-seed"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   int `gorm:"primaryKey,autoIncrement"`
	Name string
	Age  int
}

func CreateUser(db *gorm.DB, name string, age int) error {
	return db.Create(&User{Name: name, Age: age}).Error
}

var seeds = seed.Seeds{
	{
		Name: "CreateUser",
		Run: func(db *gorm.DB) error {
			return CreateUser(db, "John Doe", 30)
		},
	},
}

func main() {
	dsn := "host=host.docker.internal user=example password=example dbname=example port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = conn.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	seeder := seed.NewSeeder(seed.WithDatabase(conn))
	err = seeder.Seed(context.Background(), seeds...)
	if err != nil {
		panic(err)
	}
}
