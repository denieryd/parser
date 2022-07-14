package repo

import (
    "gorm.io/gorm"
    "parser/internal/db_models"
)

func CreateNewGoods(db *gorm.DB, goods []db_models.Goods) {
    for _, item := range goods {
        db.Create(&item)
    }
}
