package hello

import (
	"go-gin-api/internal/grpc/hello/rpc/client"
	"google.golang.org/grpc"
	"testing"
	"context"
)

func TestGetSpeak(t *testing.T) {
	conn, err := grpc.Dial(":11000", grpc.WithInsecure())
	if err != nil {
		t.Error("did not connect", err)
	}
	defer conn.Close()
	c := client.NewHelloClient(conn)
	res, err := c.GetSpeak(context.Background(), &client.GetSpeakRequest{
		Content: "hhh",
	})
	if err != nil {
		t.Error("could not hello server", err)
	}

	t.Logf("####### get server hello response: %s", res.Message)
}