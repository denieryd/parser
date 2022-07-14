package main

import (
    log "github.com/sirupsen/logrus"
    "parser/internal/config"
    "parser/internal/parsing"
    "parser/internal/repo"
    storepkg "parser/internal/storage"
    "sync"
)

func main() {
    cfg, err := config.NewConfig()
    if err != nil {
        log.Fatalf("cannot get config to start app %v", err)
    }

    storage := storepkg.GetNewStorage()
    if err = storage.Open(&cfg); err != nil {
        log.Fatalf("cannot connect to database %v", err)
    }

    wg := sync.WaitGroup{}
    urlsToParse := []string{"https://sbermarket.ru/tvoydom/c/katalogh-tvoidom/produkty-pitaniia"}
    for _, url := range urlsToParse {
        wg.Add(1)
        go func(url string, wg *sync.WaitGroup, storage *storepkg.Storage) {
            log.Infof("[START PARSING] [URL] [%v]", url)
            goods := parsing.ParsePage(url)
            repo.CreateNewGoods(storage.Db, goods)
            wg.Done()
        }(url, &wg, storage)
    }

    wg.Wait()
}
