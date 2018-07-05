package main

import (
	"context"
	"fmt"
	"log"

	"github.com/igtm/protobuf-sample/pb"
	"google.golang.org/grpc"
)

func main() {
	//sampleなのでwithInsecure
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewPrefectureServiceClient(conn)
	message := &pb.InputGetPrefectures{}
	res, err := client.GetPrefectures(context.TODO(), message)
	fmt.Printf("result:%#v \n", len(res.Prefectures))
	fmt.Printf("error::%#v \n", err)
	for _, p := range res.Prefectures {
		fmt.Printf("result:%#v \n", p)
	}
}
