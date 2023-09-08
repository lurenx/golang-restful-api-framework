package dto

type AddTaskRequest struct {
	Name    string `json:"name"`
	Status  int    `json:"status"`
	OwnerID int    `json:"owner_id"`
}
