package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go/docker-go-wb/cacher" //кешер честно взят с хабра
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	cache   *cacher.Cache
	model   ModelJson
	id      parseSrtuct
	db  *sql.DB
    err error
)

type Message struct {
	Value string
}

func main() {
	defer db.Close()
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	cache = cacher.New(5*time.Minute, 10*time.Minute)
    db, err = initStore()
    if err != nil {
        log.Fatalf("failed to initialise the store: %s", err)
    }

	e.GET("/", func(c echo.Context) error {
		return rootHandler(db, c)
	})
	e.POST("/", func(c echo.Context) error {
		var mess interface{}
		m := &Message{}

		if err := c.Bind(m); err != nil {
			log.Fatal(err)
		}
		if m.Value == "" {
			mess = "Пустой запрос"
		} else {
		    var result string
		    if val,ok := cache.Get(m.Value); ok == false {
				log.Println("В кеше отсутствует, выбор из БД")
			    result = selectMessageById(db, m.Value)
		    }else{
				log.Println("получил значение из кеша")
		        result = val.(string)
		    }
			if err := json.Unmarshal([]byte(result), &model); err != nil {
				log.Println(err)
				mess = "Не удалось выполнить запрос"
			} else {
				mess = model
			}
		}
		return c.JSON(http.StatusOK, mess)
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	var wg sync.WaitGroup
	for {
		wg.Add(1)
		go createSubscriber1(&wg)
		e.Start(":" + httpPort)
	}
	wg.Wait()

}
func createSubscriber1(wg *sync.WaitGroup) {
	defer wg.Done()
	go submitter()
}

func rootHandler(db *sql.DB, c echo.Context) error {
	return c.HTML(http.StatusOK, getHTML())
}
