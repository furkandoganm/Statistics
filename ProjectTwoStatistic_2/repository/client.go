package repository

import "go.mongodb.org/mongo-driver/mongo"

type Client struct {
	Collection *mongo.Collection
}
