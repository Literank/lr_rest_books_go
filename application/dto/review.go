/*
Package dto has all data transfer objects.
*/
package dto

type ReviewBody struct {
	BookID  uint   `json:"book_id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
