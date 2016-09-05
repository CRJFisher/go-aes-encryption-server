package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

var cipertextStrings *mgo.Collection

// Encrypted models the object that will be stored on the database
type Encrypted struct {
	ObjID            int    `json:"objId"`
	EncryptedContent string `json:"content"`
}

// StartDB connects to MongoDB and sets mode to monotonic
func StartDB() {
	var err error
	session, err = mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)

	cipertextStrings = session.DB("ciphertext-store").C("cipertextStrings")

}

// StopDB closes the session
func StopDB() {
	err := session.DB("ciphertext-store").DropDatabase()
	if err != nil {
		panic(err)
	}
	session.Close()
}

// CreateEncryptedInDB manages the ID and adds an Encrypted object to the DB
func CreateEncryptedInDB(w http.ResponseWriter, r *http.Request) {
	var data Encrypted

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &data); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//TODO: handle ID collision, change ID on obj if necessary
	err = cipertextStrings.Insert(data)
	if mgo.IsDup(err) {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

// GetEncryptedFromDB queries MongoDB for the Encrypted object by ID
func GetEncryptedFromDB(w http.ResponseWriter, r *http.Request) {
	data := Encrypted{}

	dbObjID, _ := strconv.Atoi(mux.Vars(r)["id"])

	var encrypteds []Encrypted
	err := cipertextStrings.Find(bson.M{}).All(&encrypteds)

	if err != nil {
		panic(err)
	}

	for _, enc := range encrypteds {
		if enc.ObjID == dbObjID {
			data = enc
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("ciphertext-store").C("cipertextStrings")

	index := mgo.Index{
		Key:        []string{"objId"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		log.Fatal(err)
	}
}
