package main

// Encrypted models the object that will be stored on the database
type Encrypted struct {
	ID               int
	EncryptedContent string
}

// EncryptionResponse models the data returned after content post success
type EncryptionResponse struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

// DecryptionRequest models the data received by the content retrieval request
type DecryptionRequest struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

// EncryptedData is the collection used for the database
type EncryptedData []Encrypted
