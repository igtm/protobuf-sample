package main

import (
	"net/http"

	proto1 "github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/protobuf-sample/proto"
)

func main() {
	e := echo.New()
	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/prefectures", func(c echo.Context) error {

		// DBからPrefectureのリストと取ってくる
		prefectures, _ := models.GetPrefectures()

		res := ToGetPrefecturesResponse(prefectures)
		data, _ := proto1.Marshal(&res)

		return c.Blob(http.StatusOK, "application/protobuf", data)
	})

	e.Start(":80")
}

func ToGetPrefecturesResponse(prefectures []models.Prefecture) (res proto.GetPrefecturesResponse) {
	prefecturesProto := ToPrefectures(prefectures)
	res = proto.GetPrefecturesResponse{
		Prefectures: prefecturesProto,
	}
	return
}

func ToPrefectures(prefectures []models.Prefecture) (res []*proto.Prefecture) {
	for _, pref := range prefectures {
		prefProto := ToPrefecture(&pref)
		res = append(res, &prefProto)
	}
	return
}

func ToPrefecture(pref *models.Prefecture) (chProto proto.Prefecture) {
	chProto = proto.Prefecture{
		Id:     int64(pref.Id),
		Name:   pref.Name,
		Romaji: pref.Romaji,
	}
	return
}

type Prefecture struct {
	Id     int    `json:"id"        bson:"_id"`
	Name   string `json:"name"      bson:"name"`
	Romaji string `json:"romaji"    bson:"romaji"`
}

type Prefectures struct {
	Prefectures []Prefecture `json:"prefectures"`
}
