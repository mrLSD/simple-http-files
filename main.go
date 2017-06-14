package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
)

const PATH = "/home/evgeny/d/Видео/Молодой Папа  (Season 01) BaibaKo 720/"

func main() {
	e := echo.New()
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/**/", getMainfunc)
	e.Static("/assets", PATH)
	e.Logger.Fatal(e.Start(":3040"))
}

func getMainfunc(c echo.Context) error {
	files, err := ioutil.ReadDir(PATH)
	if err != nil {
		panic(err)
	}

	type FileList struct {
		Name string
		Path string
	}
	var filesList []FileList
	for _, file := range files {
		fileData := FileList{
			Name: file.Name(),
			Path: "/assets/" + file.Name(),
		}
		filesList = append(filesList, fileData)
	}
	return c.Render(http.StatusOK, "index", filesList)
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
