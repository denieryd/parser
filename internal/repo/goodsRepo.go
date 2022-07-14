package repo

import (
    log "github.com/sirupsen/logrus"
    "gorm.io/gorm"
    "parser/internal/db_models"
)

func CreateNewGoods(db *gorm.DB, goods []db_models.Goods) {
    for _, item := range goods {
        log.Info(item.Price)
        db.Create(&item)
    }
}
