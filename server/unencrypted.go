package main

// UnencryptedRequest models the content post json
type UnencryptedRequest struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}
