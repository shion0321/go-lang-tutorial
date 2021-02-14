package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "hoge/handler"
)

func main() {
    // Echoのインスタンス作る
    e := echo.New()

    // 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
		e.Use(interceptor.BasicAuth())

    // ルーティング
    e.GET("/hello/:username", handler.MainPage(),interceptor.BasicAuth())

    // サーバー起動
    e.Start(":8080")    //ポート番号指定してね
}