package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

)

type manager struct {
	Connection 	*mongo.Client
	Ctx			context.Context
	Cancel 		context.CancelFunc
}

var Mgr manager

//db connection function

func connectDb(){

	uri := "localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s%s","mongodb://", uri)))
	
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
		return
	}

	Mgr = manager{Connection:client, Ctx:ctx, Cancel: cancel}
	fmt.Println("🚀Database Connected....!!")
}

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc){
	
	//cancelfunc to cnacel to context
	defer cancel()

	//client provies a method to close
	defer func(){
		//client.Disconnect method also has deadline
		//returns error if any
		if err := client .Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func main(){
	
	//calling connction function
	connectDb()
}