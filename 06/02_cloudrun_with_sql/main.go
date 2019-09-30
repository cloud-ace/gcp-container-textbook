package main

import (
	"context"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"

	"github.com/0Delta/CloudRunSample/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Env 環境変数から読み込む値リスト
type Env struct {
	Port string `required:"true" split_words:"true"`
}

var templates map[string]*template.Template

// Template Echoフレームワークにわたすテンプレートエンジン
type Template struct {
}

// Render テンプレート生成ロジック
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	log.Println("get template", templates)
	return templates[name].ExecuteTemplate(w, "base.html", data)
}

// init 初期化、共通テンプレートを適用して保持しておく
func init() {
	var baseTemplate = "templates/base.html"
	templates = make(map[string]*template.Template)
	//templates["index"] = template.Must(
	//  template.ParseFiles(baseTemplate, "templates/welcome.html"))
	templates["chatroom"] = template.Must(
		template.ParseFiles(baseTemplate, "templates/chatroom.html"))
	templates["chatlog"] = template.Must(
		template.ParseFiles(baseTemplate, "templates/chatlog.html"))
}

// main エントリポイント
func main() {
	e := echo.New()

	var goenv Env
	if err := envconfig.Process("", &goenv); err != nil {
		e.Logger.Fatal("Failed to Loading env.", err)
		return
	}

	t := &Template{}
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		err := c.Render(http.StatusOK, "chatroom", struct{ Name string }{Name: "ああああ"})
		if err != nil {
			log.Println(err)
		}
		return nil
	})
	e.POST("/", PostFnc)
	e.GET("/chatlog", GetLog)

	e.Logger.Fatal(e.Start(":" + goenv.Port))
}

// GetLog ログ画面を表示する
func GetLog(c echo.Context) error {
	ctx := context.Background()
	dat, err := handler.GetRecords(ctx)
	if err != nil {
		dat = []map[string]string{
			{"error": err.Error()},
		}
	}
	err = c.Render(http.StatusOK, "chatlog", struct {
		Data []map[string]string
	}{
		Data: dat,
	})
	if err != nil {
		log.Println(err)
	}
	return nil
}

// PostFnc POSTリクエストが来たら来たデータをDBに格納してトップ画面を表示する
func PostFnc(c echo.Context) error {
	err := handler.AddRecords(c)
	if err != nil {
		c.String(http.StatusBadRequest, "Error : "+err.Error())
		return err
	}
	d := new(struct{ Name string })
	err = c.Bind(d)
	if err != nil {
		log.Println(err)
		return err
	}
	err = c.Render(http.StatusOK, "chatroom", struct{ Name string }{Name: d.Name})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
