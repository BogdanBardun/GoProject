package database

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"math/rand"
)
var ProductsDB *mgo.Database
var Session *mgo.Session
type Product struct{
	Id bson.ObjectId `bson:"_id"`
	Part string `bson:"part"`
	Company string `bson:"company"`
	Price int `bson:"price"`
	Quantity int `bson:"quantity"`
	Category string `bson:"category"`
}
type User struct{
	Id bson.ObjectId `bson:"_id"`
	Username string `bson:"username"`
	Login string `bson:"login"`
	Password string `bson:"password"`
}
func Init() {
	Session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		panic(err)
	}
	ProductsDB = Session.DB("GeneralDB")
	productsCollection := Session.DB("GeneralDB").C("Products")
	p1 := &Product{Id:bson.NewObjectId(), Part:"WH1258", Company:"Continental", Price:450, Quantity:20, Category:"Chassis"}
	p2 := &Product{Id:bson.NewObjectId(), Part:"TR3344", Company:"Brabus", Price:1200, Quantity:5, Category:"Chassis"}
	p3 := &Product{Id:bson.NewObjectId(), Part:"B55", Company:"AMG", Price:600, Quantity:11, Category:"Body"}
	p4 := &Product{Id:bson.NewObjectId(), Part:"S44", Company:"Alcantara", Price:750, Quantity:50, Category:"Interior"}
	p5 := &Product{Id:bson.NewObjectId(), Part:"N55T", Company:"GM", Price:50, Quantity:40, Category:"Engine"}
	p6 := &Product{Id:bson.NewObjectId(), Part:"VV1", Company:"MG", Price:1100, Quantity:3, Category:"Engine"}
	p7 := &Product{Id:bson.NewObjectId(), Part:"NG4", Company:"BMW", Price:800, Quantity:8, Category:"Body"}
	p8 := &Product{Id:bson.NewObjectId(), Part:"JK122", Company:"Brembo", Price:500, Quantity:30, Category:"Interior"}
	err = productsCollection.Insert(p1, p2, p3, p4, p5, p6, p7, p8)
	if err != nil{
		fmt.Println(err)
	}
	mnmspc := [15]string{"Continental", "Brabus", "AMG", "Alcantara", "GM", "MG", "BMW", "Brembo", "MB", "Ford", "Renault", "Alpine", "Bridgestone", "Michelin", "Chevrolet"}
	cnmspc := [4]string{"Engine", "Interior", "Chassis", "Body"}
	for i := 0; i < 100; i++ {
		pp := &Product{Id:bson.NewObjectId(), Part: strconv.Itoa(rand.Int() % 1000), Company:mnmspc[rand.Int() % 15], Price:rand.Int() % 10000, Quantity:rand.Int() % 1000, Category:cnmspc[rand.Int() % 4]}
		err = productsCollection.Insert(pp)
		if err != nil{
			fmt.Println(err)
		}
	}
  usersCollection := Session.DB("GeneralDB").C("Users")
	u1 := &User{Id:bson.NewObjectId(), Username:"Joe", Login:"whosjoe", Password:"Joemama"}
	u2 := &User{Id:bson.NewObjectId(), Username:"Boban", Login:"BobanMarjanovic", Password:"srbijastrong"}
	u3 := &User{Id:bson.NewObjectId(), Username:"Jeoff", Login:"Geff", Password:"geoffe"}
	err = usersCollection.Insert(u1, u2, u3)
	if err != nil{
		fmt.Println(err)
	}
}
