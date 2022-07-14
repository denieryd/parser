package config

import (
    "github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
    Database struct {
        Host     string `json:"host" env-description:"Db host" env-default:"localhost"`
        Port     string `json:"port" env-description:"Db port" env-default:"5432"`
        DbName   string `json:"db_name" env-description:"Db name" env-default:"dbname"`
        SSLMode  string `json:"ssl_mode" env-description:"Db ssl mode" env-default:"disable"`
        User     string `json:"user" env-description:"Db user" env-default:"postgres"`
        Password string `json:"password" env-description:"Db password" env-default:"pass"`
    } `json:"database"`
    App struct {
        BindAddr string `json:"bind_addr" env-description:"Server bind addr" env-default:":8080"`
        LogLevel string `json:"log_level" env-description:"App log level" env-default:"info"`
    } `json:"app"`
}

func NewConfig() (cfg Config, _ error) {
    return cfg, cleanenv.ReadConfig("./configs/config.json", &cfg)
}
