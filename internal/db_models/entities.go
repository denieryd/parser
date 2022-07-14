package db_models

import (
    "github.com/google/uuid"
)

type Goods struct {
    ID       uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
    Name     string    `gorm:"column:name"`
    URL      string    `gorm:"column:url"`
    URLImage string    `gorm:"column:url_image"`
    Price    float32   `gorm:"column:price"`
}
