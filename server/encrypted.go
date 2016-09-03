package main

// Encrypted models the object that will be stored on the database
type Encrypted struct {
	ID               int
	EncryptedContent []byte
	Key              []byte
}

// EncryptionResponse models the data returned after content post success
type EncryptionResponse struct {
	ID  int    `json:"id"`
	Key []byte `json:"key"`
}

// EncryptedData is the collection used for the database
type EncryptedData []Encrypted
