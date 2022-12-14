package main

import (
	"context"
	"fmt"
	pb "github.com/shamanis/birdie/internal/app/birdie"
	"github.com/shamanis/birdie/internal/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"time"
)

var logger = logging.New()

func main() {
	go pb.StartService()
	time.Sleep(1 * time.Second) // Wait a server start

	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("127.0.0.1:50051", opts...)
	if err != nil {
		logger.Panicf("fail dialing: %s", err)
	}
	defer conn.Close()

	client := pb.NewBirdieClient(conn)

	fmt.Println("Store entry key=key1, value=value1")
	startStore := time.Now()
	res, err := client.Store(context.Background(), &pb.Entry{Key: "key1", Value: []byte("value1"), Ttl: 5})
	if err != nil {
		logger.Panicf("fail store: %s", err)
	}
	fmt.Printf("Response: status=%d time=%f\n", res.Status, time.Since(startStore).Seconds())

	fmt.Println("Load entry key=key1")
	startLoad := time.Now()
	entry, err := client.Load(context.Background(), &pb.LoadQuery{Key: "key1"})
	if err != nil {
		logger.Panicf("fail load: %s", err)
	}
	fmt.Printf("Entry value: %s time=%f\n", entry.Value, time.Since(startLoad).Seconds())

	fmt.Println("Load entry key=not-found")
	_, err = client.Load(context.Background(), &pb.LoadQuery{Key: "not-found"})
	if err != nil {
		st, _ := status.FromError(err)
		fmt.Printf("Status: [%d] %s\n", st.Code(), st.Message())
	}
}
