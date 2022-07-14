package db_models

import (
    "github.com/google/uuid"
)

type dbModel struct{}

type Goods struct {
    dbModel
    ID       uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
    Name     string    `gorm:"column:name;unique"`
    URL      string    `gorm:"column:url;not null;unique"`
    URLImage string    `gorm:"column:url_image;not null;unique"`
    Price    float32   `gorm:"column:price;not null;unique"`
}
