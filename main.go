package main

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/igtm/protobuf-sample/sample"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/prefectures", func(c echo.Context) error {

		// DBからPrefectureのリストと取ってくる
		res := sample.GetPrefecturesResponse{
			Prefectures: []*sample.Prefecture{
				&sample.Prefecture{
					Id:     1,
					Name:   "北海道",
					Romaji: "hokkaido",
				},
				&sample.Prefecture{
					Id:     47,
					Name:   "沖縄県",
					Romaji: "okinawa",
				},
			},
		}
		data, _ := proto.Marshal(&res)

		return c.Blob(http.StatusOK, "application/protobuf", data)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
