package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var currentID int

var dataObjs EncryptedData

var url = "http://127.0.0.1:8082"

// RepoFindObj gets the object with given ID from the database
func RepoFindObj(ID int) Encrypted {
	var data Encrypted

	if UseRemoteStore {

		client := &http.Client{}
		req, err := http.NewRequest("GET", url+"/retrieve/"+strconv.Itoa(ID), nil)
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &data); err != nil {
			panic(err)
		}

		return data

	}

	for _, enc := range dataObjs {
		if enc.ID == ID {
			return enc
		}
	}

	// return empty Obj if not found
	return Encrypted{}
}

// RepoCreateObj manages the ID of the object and then stores the object
func RepoCreateObj(enc Encrypted) Encrypted {
	var data Encrypted

	if UseRemoteStore {

		client := &http.Client{}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(enc)
		req, err := http.NewRequest("POST", url+"/store", b)
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &data); err != nil {
			panic(err)
		}

		return data

	}

	//TODO: If the ID already exists, change it so it doesn't collide
	//----> Implement idAlreadyExists

	if enc.ID == 0 {
		currentID++
		enc.ID = currentID
	}

	dataObjs = append(dataObjs, enc)
	return enc
}

func idAlreadyExists(id int) bool {
	for _, enc := range dataObjs {
		if enc.ID == id {
			return true
		}
	}
	return false
}
