package main

import (
	"aodebiao/ent"
	"aodebiao/model"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	model.InitDb()
	defer model.Client.Close()

	CreateUser(context.TODO(),model.Client)
}
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	//user, err := client.User.Create().SetCreateAt(time.Now()).SetAvatarURL("aodebiao").SetUserName("aodebiao").SetPassword("aodebiao").SetDelFlag(1).SetID("adadasd").Save(ctx)
	user, err := client.User.Create().SetCreatedAt(time.Now()).SetAvatarURL("aodebiao1").SetUserName("aodebiao1").SetPassword("aodebiao1").SetDelFlag(int(1)).SetID("1231").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user :%v ", err)
	}
	return user, nil
}
