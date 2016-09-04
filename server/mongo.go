package main

import (
	"fmt"
	"strconv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

var cipertextStrings *mgo.Collection

// StartDB connects to MongoDB and sets mode to monotonic
func StartDB() {
	var err error
	session, err = mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	err = session.DB("ciphertext-store").DropDatabase()
	if err != nil {
		panic(err)
	}

	cipertextStrings = session.DB("ciphertext-store").C("cipertextStrings")

}

// StopDB closes the session
func StopDB() {
	session.Close()
}

// GetEncryptedFromDB queries MongoDB for the Encrypted object by ID
func GetEncryptedFromDB(ID int) Encrypted {
	fmt.Println("get encrypted from DB")
	var encrypteds []Encrypted
	err := cipertextStrings.Find(bson.M{}).All(&encrypteds)
	// result := Encrypted{}
	// err := cipertextStrings.Find(bson.M{"ID": ID}).One(&result)
	if err != nil {
		panic(err)
	}
	// fmt.Println("result ID: " + strconv.Itoa(result.ID))
	for _, enc := range encrypteds {
		fmt.Println("encrypted id: " + strconv.Itoa(enc.ID))
		if enc.ID == ID {
			return enc
		}
	}

	return Encrypted{}
}

// CreateEncryptedInDB manages the ID and adds an Encrypted object to the DB
func CreateEncryptedInDB(enc Encrypted) Encrypted {

	//TODO: handle ID collision
	cipertextStrings.Insert(enc)

	return enc
}
