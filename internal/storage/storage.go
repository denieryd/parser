package storage

import (
    "fmt"
    _ "github.com/lib/pq"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "parser/internal/config"
    "parser/internal/db_models"
)

func GetNewStorage() *Storage {
    return &Storage{}
}

type Storage struct {
    Db *gorm.DB
}

func (s *Storage) Open(cfg *config.Config) error {
    dsn := fmt.Sprintf("host=%s dbname=%s sslmode=%s user=%s password=%s port=%s",
        cfg.Database.Host, cfg.Database.DbName, cfg.Database.SSLMode,
        cfg.Database.User, cfg.Database.Password, cfg.Database.Port)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    if err := db.AutoMigrate(&db_models.Goods{}); err != nil {

    }

    s.Db = db
    return nil
}
