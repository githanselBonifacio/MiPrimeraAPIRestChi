package models

import (
	"errors"
	"time"
)

const maxLenghtInComments = 400

//Review represent an anon review from some website

type Review struct {
	Id      int64
	Stars   int       //1-5
	Comment string    //max 400 chars
	Date    time.Time //created at
}

//CreateReview commad to create a new review
type CreateReviewCmd struct {
	Stars   int    `json:"stars"`
	Comment string `json:"comment"`
}

//method create review
func (cmd *CreateReviewCmd) validate() error { //create an pointer
	if cmd.Stars < 1 || cmd.Stars > 5 {
		return errors.New("Stars must be between 1-5")
	}
	if len(cmd.Comment) > maxLenghtInComments {
		return errors.New("Comment must be less that 400  charts")
	}
	return nil
}
