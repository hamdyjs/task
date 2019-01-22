package cmd

type task struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Status int    `json:"status"`
}
