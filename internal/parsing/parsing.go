package parsing

import (
    "github.com/antchfx/htmlquery"
    "github.com/google/uuid"
    log "github.com/sirupsen/logrus"
    "golang.org/x/net/html"
    "parser/internal/db_models"
    "regexp"
    "strconv"
)

func ParsePage(urlToParse string) []db_models.Goods {
    doc, err := htmlquery.LoadURL(urlToParse)
    if err != nil {
        log.Warnf("Cannot parse given url %v", urlToParse)
        return []db_models.Goods{}
    }

    productPage := htmlquery.Find(doc, "//div[@class='ProductsGrid_styles_grid__zigKP']")
    if productPage == nil {
        return []db_models.Goods{}
    }

    productCards := htmlquery.Find(productPage[0], "//div[@class='ProductCard_styles_root__uKQJb LargePromoProduct_styles_root__HgnyZ']")
    goods := parseProductCards(productCards)

    return goods
}

func parseProductCards(productCards []*html.Node) []db_models.Goods {
    resultedGoods := make([]db_models.Goods, 0)
    priceRegExp := regexp.MustCompile(`\d{2},\d{2}`)

    for _, productCard := range productCards {
        id, _ := uuid.NewUUID()
        goodInstance := db_models.Goods{ID: id}

        aElem := htmlquery.FindOne(productCard, "//a")
        if aElem != nil {
            link := htmlquery.SelectAttr(aElem, "href")
            goodInstance.URL = "sbermarket.ru" + link
        }

        pictureElem := htmlquery.FindOne(aElem, "//picture[@class='Picture_root__qwGYp ProductCard_styles_picture__6X2V4']")
        if pictureElem != nil {
            imageElem := htmlquery.FindOne(pictureElem, "//img[@class='Image_root__QyHLt ProductCard_styles_image__0XgSg']")
            urlImage := htmlquery.SelectAttr(imageElem, "src")
            goodInstance.URLImage = urlImage
        }

        nameElem := htmlquery.FindOne(aElem, "//h3[@class='ProductCard_styles_title__vb8ha']")
        if nameElem != nil {
            goodInstance.Name = htmlquery.InnerText(nameElem)
        }

        priceBlock := htmlquery.FindOne(aElem, "//div[@class='ProductCardPrice_styles_root__2QeEZ ProductCard_styles_price__huT9y']")
        if priceBlock != nil {
            priceElem := htmlquery.FindOne(aElem, "//div[@class='ProductCardPrice_styles_price__n51qG ProductCardPrice_styles_accent__jQFs_']")
            if priceElem != nil {
                text := htmlquery.InnerText(priceElem)
                priceStr := priceRegExp.FindStringSubmatch(text)[0]
                price, err := strconv.ParseFloat(priceStr, 64)
                if err != nil {
                    goodInstance.Price = float32(price)
                }
            }
        }

        resultedGoods = append(resultedGoods, goodInstance)
    }

    return resultedGoods
}
