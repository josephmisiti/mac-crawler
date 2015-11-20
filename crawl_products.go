package main

import(
    "fmt"
    "log"
    
    "github.com/PuerkitoBio/goquery"
)

func getUrls() []string {
    // urls := []string{
    //     "http://www.maccosmetics.com/products/13834/Products/Makeup/Eyes/Brow",
    //     "http://www.maccosmetics.com/products/13840/Products/Makeup/Eyes/Shadow",
    //     "http://www.maccosmetics.com/products/13835/Products/Makeup/Eyes/Eye-Palettes-Kits",
    //     "http://www.maccosmetics.com/products/14756/Products/Makeup/Eyes/Eye-Primer",
    //     "http://www.maccosmetics.com/products/13837/Products/Makeup/Eyes/Lash",
    //     "http://www.maccosmetics.com/products/13838/Products/Makeup/Eyes/Liner",
    //     "http://www.maccosmetics.com/products/13839/Products/Makeup/Eyes/Mascara",
    //     "http://www.maccosmetics.com/products/13825/Products/Skincare/Primers",
    //     "http://www.maccosmetics.com/products/14771/Products/Skincare/BB-CC",
    //     "http://www.maccosmetics.com/products/13824/Products/Skincare/Moisturizers",
    //     "http://www.maccosmetics.com/products/13825/Products/Skincare/Primers",
    //     "http://www.maccosmetics.com/products/13826/Products/Skincare/Removers",
    //     "http://www.maccosmetics.com/products/13827/Products/Skincare/Travel-Size",
    //     "http://www.maccosmetics.com/products/13804/Products/Brushes-Tools/Brushes/Eye-Brushes",
    //     "http://www.maccosmetics.com/products/13802/Products/Brushes-Tools/Brushes/All-Brushes",
    //     "http://www.maccosmetics.com/products/13803/Products/Brushes-Tools/Brushes/Brush-Kits",
    //     "http://www.maccosmetics.com/products/13805/Products/Brushes-Tools/Brushes/Face-Brushes",
    //     "http://www.maccosmetics.com/products/13806/Products/Brushes-Tools/Brushes/Lip-Brushes",
    //     "http://www.maccosmetics.com/products/13815/Products/Fragrance/Turquatic",
    // }
    
    urls := []string{
        "http://www.maccosmetics.com/products/13834/Products/Makeup/Eyes/Brow",
        // "http://www.maccosmetics.com/products/13840/Products/Makeup/Eyes/Shadow",
        // "http://www.maccosmetics.com/products/13835/Products/Makeup/Eyes/Eye-Palettes-Kits",
        // "http://www.maccosmetics.com/products/14756/Products/Makeup/Eyes/Eye-Primer",
        // "http://www.maccosmetics.com/products/13837/Products/Makeup/Eyes/Lash",
        // "http://www.maccosmetics.com/products/13838/Products/Makeup/Eyes/Liner",
        // "http://www.maccosmetics.com/products/13839/Products/Makeup/Eyes/Mascara",
        // "http://www.maccosmetics.com/products/13825/Products/Skincare/Primers",
        // "http://www.maccosmetics.com/products/14771/Products/Skincare/BB-CC",
        // "http://www.maccosmetics.com/products/13824/Products/Skincare/Moisturizers",
        // "http://www.maccosmetics.com/products/13825/Products/Skincare/Primers",
        // "http://www.maccosmetics.com/products/13826/Products/Skincare/Removers",
        // "http://www.maccosmetics.com/products/13827/Products/Skincare/Travel-Size",
        // "http://www.maccosmetics.com/products/13804/Products/Brushes-Tools/Brushes/Eye-Brushes",
        // "http://www.maccosmetics.com/products/13802/Products/Brushes-Tools/Brushes/All-Brushes",
        // "http://www.maccosmetics.com/products/13803/Products/Brushes-Tools/Brushes/Brush-Kits",
        // "http://www.maccosmetics.com/products/13805/Products/Brushes-Tools/Brushes/Face-Brushes",
        // "http://www.maccosmetics.com/products/13806/Products/Brushes-Tools/Brushes/Lip-Brushes",
        // "http://www.maccosmetics.com/products/13815/Products/Fragrance/Turquatic",
    }
    return urls
}

func parseColors(s *goquery.Selection) string {
    colors := ""
    s.Each(func(i int, s *goquery.Selection) {
        colors += s.Text()
    })
    return colors
}

func parseWeight(s *goquery.Selection) string { 
    //colors := ""
    return s.Text()
}

func crawlProducts() {
    urls := getUrls()
    for i := 0; i<len(urls); i++ {
        
        doc, err := goquery.NewDocument(urls[i]) 
          if err != nil {
            log.Fatal(err)
        }
        
        doc.Find("div.grid--mpp__item").Each(func(i int, s *goquery.Selection) {
            // name := s.Find(".product__name").Text()
            // description := s.Find(".product__description-short").Text()
            // price := s.Find(".product__price").Text()
            // colors := parseColors(s.Find(".shade-picker__color-name"))
           weight := parseWeight(s.Find(".product__size"))
            
            fmt.Printf("%s \n", weight)
            //fmt.Printf("%s %s %s %s %s\n", name, description, price, colors, weight)
        })
        
        
    }
}

func main() {
    fmt.Println("crawling www.mac.com..... \n")
    crawlProducts()
}

