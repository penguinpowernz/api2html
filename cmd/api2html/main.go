package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "-c", "config.json", "the config file to use")
	flag.Parse()

	cfg := config{}
	data, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}

	api := gin.Default()
	api.LoadHTMLGlob("./templates/*")

	for _, p := range cfg.Pages {
		func(p page) {
			log.Printf("Setting up page: %s", p.URI)
			api.GET(p.URI, func(c *gin.Context) {
				url := p.APIURL
				for _, prm := range c.Params {
					url = strings.ReplaceAll(url, ":"+prm.Key, prm.Value)
				}

				var view interface{}
				r, err := http.Get(url)
				log.Printf("calling %s API for %s", p.APIURL, p.URI)
				if err != nil {
					c.AbortWithError(503, err)
					return
				}

				defer r.Body.Close()
				data, err := ioutil.ReadAll(r.Body)
				err = json.Unmarshal(data, &view)

				c.Header("Cache-Control", "private max-age="+strconv.Itoa(p.CacheExpiry))

				c.HTML(r.StatusCode, p.Template+".html", view)
			})
		}(p)
	}

	api.Run()

}

type config struct {
	Pages []page
}

type page struct {
	APIURL      string `json:"api_url"`
	URI         string `json:"uri"`
	Template    string `json:"template"`
	CacheExpiry int    `json:"cache_expiry"`
}
