package main

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

type Proxy struct {
    ID     int
    Type   string
    Secret string
    Link   string
}

var proxies []Proxy
var proxyID = 1

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

    serverIP := os.Getenv("SERVER_IP")
    if serverIP == "" {
        serverIP = "127.0.0.1"
    }

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
            "tg://proxy?server=%s&port=443&secret=%s",
            serverIP,
            secret,
        )

        proxies = append(proxies, Proxy{
            ID:     proxyID,
            Type:   ptype,
            Secret: secret,
            Link:   link,
        })

        proxyID++

        c.Redirect(http.StatusFound, "/")
    })

    r.POST("/delete/:id", func(c *gin.Context) {
        id := c.Param("id")

        newList := []Proxy{}

        for _, p := range proxies {
            if fmt.Sprintf("%d", p.ID) != id {
                newList = append(newList, p)
            }
        }

        proxies = newList

        c.Redirect(http.StatusFound, "/")
    })

    r.Run(":8080")
}
