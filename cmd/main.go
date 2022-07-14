package main

import (
    "fmt"
    "github.com/antchfx/htmlquery"
    log "github.com/sirupsen/logrus"
)

func main() {
    //cfg, err := config.NewConfig()
    //if err != nil {
    //    log.Fatalf("cannot get config to start app %v", err)
    //}
    //
    //storage := storepkg.GetNewStorage()
    //if err = storage.Open(&cfg); err != nil {
    //    log.Fatalf("cannot connect to database %v", err)
    //}

    urlToParse := "https://www.ozon.ru/category/moloko-9283/"

    doc, err := htmlquery.LoadURL(urlToParse)
    if err != nil {
        panic(err)
    }

    nodes := htmlquery.Find(doc, "//a[@href]")
    if err != nil {
        panic(err)
    }

    log.Info(nodes)

    for _, n := range nodes {
        fmt.Println(htmlquery.SelectAttr(n, "href")) // output @href value
    }
}
