package repository

import (
	"classroom/model"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	SetGroupHash(group *model.Group) error 
	SetGroup(group *model.Group) error 
	GetGroup(id string) (model.Group, error) 
}

type cache struct {
	client *redis.Client
}

func NewCache() Cache {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	addr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       1,
	})

	ctx := context.Background()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	cache := &cache{client: client}

	return cache
}

func (c *cache) SetGroup(group *model.Group) error {
	ctx := context.Background()

	key := fmt.Sprintf("cod_group:%v", group.ID)
	groupMarshal, err := group.Value()
	if err != nil {
		return err
	}

	err = c.client.Set(ctx, key, groupMarshal, 60*60).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *cache) GetGroup(id string) (model.Group, error) {
	ctx := context.Background()
	key := fmt.Sprintf("cod_group:%v", id)
	var group model.Group

	err := c.client.Get(ctx, key).Scan(&group)
	if err != nil {
		return group, err
	}

	return group, nil
}

func (c *cache) SetGroupHash(group *model.Group) error {
	// ctx := context.Background()
	
	// key := fmt.Sprintf("cod_group:%v", group.ID)
	// classesMarshal, err := group.Classes.Value()
	// group.Classes = classesMarshal
	// if err != nil {
	// 	return err
	// }
	// err = c.client.HSet(ctx, key, groupMarshal).Err()
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (c *cache) GetGroupHash(key string, group *model.Group) (*model.Group, error) {
	ctx := context.Background()

	res := c.client.HGetAll(ctx, key)

	err := res.Scan(&group)
	if err != nil {
		return &model.Group{}, err
	}

	err = res.Scan(&group.Classes)
	if err != nil {
		return &model.Group{}, err
	}

	return group, nil
}
