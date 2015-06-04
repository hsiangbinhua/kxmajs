package utils

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"com.kxfax/model"
	"fmt"
)

func InsertDB(majs model.Majs) {
	//session, err := mgo.Dial("mongodb://root:123456@10.170.22.51:27017")
	session, err := mgo.Dial("mongodb://root:123456@192.168.1.200:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("admin").C("majs")
	err = c.Insert(majs)
	if err != nil {
		log.Fatal(err)
	}
	result := model.Majs{}
	err = c.Find(bson.M{"u_account": "10728"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", result)
}
