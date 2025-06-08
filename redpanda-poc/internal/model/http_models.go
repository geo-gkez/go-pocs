package model

type ProduceMessageRequest struct {
	Key     string `json:"key" binding:"required"`
	Message string `json:"message" binding:"required"`
}
