package models

type Mail struct {
	Id      string `json:"id" firestore:"id"`
	To      string `json:"to" firestore:"to"`
	Subject string `json:"subject" firestore:"subject"`
	Body    string `json:"body" firestore:"body"`
}
