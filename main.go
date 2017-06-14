package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
)

const PATH = "/home/evgeny/d/Видео/Молодой Папа  (Season 01) BaibaKo 720/"

func main() {
	files, err := ioutil.ReadDir(PATH)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/**/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.Static("/assets/", PATH)

	// Start server
	e.Logger.Fatal(e.Start(":304" +
		"0"))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
