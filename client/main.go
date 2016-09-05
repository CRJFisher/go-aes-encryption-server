package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type unencryptedData struct {
	ID                 int `json:"id"`
	UnencryptedContent string `json:"content"`
}

type encryptionResponse struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

var url = "http://127.0.0.1:8080"

func main() {

	testID := "1234"
	testContent := "test content"

	key, err := Store([]byte(testID), []byte(testContent))
	if err != nil {
		panic(err)
	}

	plainText, err := Retrieve([]byte("1234"), key)

	if string(plainText[:]) == testContent {
		fmt.Println("encrypt->decrypt successful!")
	}

}

// Store makes a request to the encryption server and returns the generated key
func Store(id, payload []byte) (aesKey []byte, err error) {

	idString := string(id[:4])
	payloadString := string(payload)

	idInt, _ := strconv.Atoi(idString)
	reqData := unencryptedData{idInt, payloadString}

	client := &http.Client{}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(reqData)
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

	respData := encryptionResponse{}
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &respData)

	return []byte(respData.Key), err
}

// Retrieve passes the key to the encryption server and returns the original message
func Retrieve(id, aesKey []byte) (payload []byte, err error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url+"/retrieve", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
    q.Add("Key", string(aesKey[:]))
    q.Add("ID", string(id[:4]))
    req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	respData := unencryptedData{}

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &respData)

	return []byte(respData.UnencryptedContent), err
}
