package main

import (
	ent "aodebiao/ent"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	client, err := ent.Open("mysql", "root:1460088689@tcp(localhost:3306)/aodebiao?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	//CreateUser(context.TODO(), client)

	fmt.Print(time.Now())
}
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	user, err := client.User.Create().SetCreateAt(time.Now()).SetAvatarURL("1231").SetUserName("loocc").SetPassword("1231").SetDelFlag(1).SetID("12313").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user :%v ", err)
	}
	return user, nil
}
