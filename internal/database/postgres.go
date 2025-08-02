package database

import (
	"log"
	"os"

    "go-fiber-wire/internal/features/user"
    "go-fiber-wire/internal/features/role"

	"gorm.io/gorm"
    "gorm.io/driver/postgres"

)

func Connect() *gorm.DB {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL is not set")
    }

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    if os.Getenv("ENV") != "production" {
	    db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

		db.AutoMigrate(
			&role.Role{},
			&user.User{},
		)
    }

    MigrateAndSeed(db)
    return db
}