package models

type InsertPostReq struct {
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Label      []string `json:"label"`
	Categories []string `json:"categories"`
}
