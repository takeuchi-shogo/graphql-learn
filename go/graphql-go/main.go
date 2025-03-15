package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/handler"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/loader"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/repository"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/schema"
	"github.com/takeuchi-shogo/graphql-learn/go/graphql-go/service/user"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("1M"))

	s, err := schema.String()
	if err != nil {
		log.Fatal(err)
	}

	schemaHandler, err := graphql.ParseSchema(s, &handler.Query{})
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository()
	userService := user.NewGetUserService(userRepository)

	e.POST("/graphql", func(c echo.Context) error {
		h := &relay.Handler{Schema: schemaHandler}
		w := httptest.NewRecorder()

		// 新しいコンテキストを作成
		loaderCollection := loader.NewLoaderCollection()
		ctx := c.Request().Context()
		// まずLoaderをアタッチ
		ctx = loaderCollection.Attach(ctx)
		// その後UserServiceを追加
		ctx = context.WithValue(ctx, loader.UserServiceKey, userService)

		// 修正したコンテキストでリクエスト
		h.ServeHTTP(w, c.Request().WithContext(ctx))
		return c.JSONBlob(http.StatusOK, w.Body.Bytes())
	})

	e.GET("/query", func(c echo.Context) error {
		s, err := schema.String()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, s)
	})

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	e.Start(":8080")
}
