package main

import (
	"fmt"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/igtm/protobuf-sample/pb"
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
		res := pb.Prefectures{
			Prefectures: []*pb.Prefecture{
				&pb.Prefecture{
					Id:     1,
					Name:   "北海道",
					Romaji: "hokkaido",
				},
				&pb.Prefecture{
					Id:     47,
					Name:   "沖縄県",
					Romaji: "okinawa",
				},
			},
		}
		data, _ := proto.Marshal(&res)
		fmt.Printf(`%v`, data)

		return c.Blob(http.StatusOK, "application/protobuf", data)
	})
	e.GET("/prefectures-json", func(c echo.Context) error {

		// DBからPrefectureのリストと取ってくる
		res := []map[string]interface{}{
            map[string]interface{}{
                "Id":     1,
                "Name":   "北海道",
                "Romaji": "hokkaido",
            },
            map[string]interface{}{
                "Id":     47,
                "Name":   "沖縄県",
                "Romaji": "okinawa",
            },
		}

		return c.JSON(http.StatusOK, res)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
