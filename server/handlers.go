package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

// ObjRetrieve takes an ID and key then returns the decrypted plaintext
func ObjRetrieve(w http.ResponseWriter, r *http.Request) {
	i, err := strconv.Atoi(r.URL.Query().Get("ID"))
	if err != nil {
		panic(err)
	}
	credential := DecryptionRequest{ID: i, Key: r.URL.Query().Get("Key")}

	dataFromStore := RepoFindObj(credential.ID)

	decryptedData, err := Decrypt(dataFromStore.EncryptedContent, []byte(credential.Key))
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(UnencryptedRequest{i, decryptedData}); err != nil {
		panic(err)
	}
}

// ObjCreate takes an ID and plaintext, then encrypts and stores it in the database
func ObjCreate(w http.ResponseWriter, r *http.Request) {
	var data UnencryptedRequest

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

	//Encrypt the data
	ciphertext, key, _ := Encrypt([]byte(data.Content))
	encrypted := Encrypted{data.ID, ciphertext}

	//Store it
	RepoCreateObj(encrypted)

	//Return the key and reference ID
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(EncryptionResponse{encrypted.ID, key}); err != nil {
		panic(err)
	}
}
