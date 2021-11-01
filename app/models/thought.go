package models

type Thought struct {
	ThoughtId string `json:"thought_id"`
	UserId    string `json:"user_id"`
	Thought   string `json:"thought"`
	IsPrivate int    `json:"is_private"`
	CreatedAt string `json:"created_at"`
}