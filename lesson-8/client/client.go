package main

import (
	pb "client/api/proto"
	"context"
	"log"

	"google.golang.org/grpc"
)

func NewItemRepositoryClient(addr string) (pb.ItemRepositoryClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return pb.NewItemRepositoryClient(conn), nil
}

func main() {
	itemRepository, err := NewItemRepositoryClient("localhost:9094")
	if err != nil {
		log.Fatal(err)
	}

	createItemReq := &pb.CreateItemRequest{
		Name:  "test_item_1",
		Price: 1400,
	}
	item, err := itemRepository.CreateItem(context.Background(), createItemReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*item)

	updateItemReq := &pb.UpdateItemRequest{
		Id:    10,
		Name:  "test_item_upd",
		Price: 1500,
	}
	item, err = itemRepository.UpdateItem(context.Background(), updateItemReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*item)
}
