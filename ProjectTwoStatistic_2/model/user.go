package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserList struct {
	Users []User `json:"users"`
}

type User struct {
	Id          primitive.ObjectID `json:"id"`
	Username    string             `json:"username"`
	Password    string             `json:"password"`
	EMail       string             `json:"eMail"`
	Status      string             `json:"status"`
	Database    string             `json:"database"`
	Collections []string           `json:"collections"`
	Videos      []string           `json:"videos"`
}

type MinUser struct {
	Id       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}

/*
go get github.com/gocolly/colly/v2
go get -u github.com/go-colly/colly
go get -u github.com/go-colly/colly/...
go get github.com/PuerkitoBio/goquery
go get github.com/dgrijalva/jwt-go
--go get go.mongodb.org/mongo-driver/bsoncls
--go get go.mongodb.org/mongo-driver/bson
--go get  go.mongodb.org/mongo-driver/x/mongo/driver/ocsp
--go get go.mongodb.org/mongo-driver/mongo
--go get golang.org/x/sync/errgroup
--go get go.mongodb.org/mongo-driver/bson/primitive
--go get -u golang.org/x/oauth2/...
--go get -u google.golang.org/api/youtube/v3
*/
