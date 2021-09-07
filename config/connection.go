package config

import (
	"log"
	"membership/pb"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connection struct{}

func NewConn() *Connection {
	return &Connection{}
}

func (c *Connection) getConn(port string) (conn *grpc.ClientConn) {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return
}

func (c *Connection) ProductConn() (client pb.ProductServiceClient) {
	port := os.Getenv("PRODUCT_RPC_PORT")
	conn := c.getConn(port)
	client = pb.NewProductServiceClient(conn)

	return
}
