package main

import (
	"fmt"
	"context"
	"log"
	"net"

	"github.com/igtm/protobuf-sample/pb"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	prefService := &PrefectureService{}
	// 実行したい実処理をseverに登録する
	pb.RegisterPrefectureServiceServer(server, prefService)
	server.Serve(listenPort)
}

type PrefectureService struct {
}

func (s *PrefectureService) GetPrefectures(ctx context.Context, message *pb.InputGetPrefectures) (*pb.Prefectures, error) {
	fmt.Println("called!")
	return &pb.Prefectures{
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
	}, nil
}
