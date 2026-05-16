package main

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

type Proxy struct {
    Type   string
    Secret string
    Link   string
}

var proxies []Proxy

func randomHex(n int) string {
    b := make([]byte, n)
    rand.Read(b)
    return hex.EncodeToString(b)
}

func generateClassic() string {
    return randomHex(16)
}

func generateDD() string {
    return "dd" + randomHex(16)
}

func generateEE(domain string) string {
    secret := randomHex(16)
    domainHex := hex.EncodeToString([]byte(domain))
    return "ee" + secret + domainHex
}

func main() {
    r := gin.Default()

    r.LoadHTMLGlob("web/templates/*")

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "proxies": proxies,
        })
    })

    r.POST("/generate", func(c *gin.Context) {
        ptype := c.PostForm("type")
        domain := c.PostForm("domain")

        var secret string

        switch ptype {
        case "classic":
            secret = generateClassic()
        case "dd":
            secret = generateDD()
        case "ee":
            if domain == "" {
                domain = "google.com"
            }
            secret = generateEE(domain)
        default:
            secret = generateClassic()
        }

        link := fmt.Sprintf(
            "tg://proxy?server=127.0.0.1&port=443&secret=%s",
            secret,
        )

        proxies = append(proxies, Proxy{
            Type:   ptype,
            Secret: secret,
            Link:   link,
        })

        c.Redirect(http.StatusFound, "/")
    })

    r.Run(":8080")
}
