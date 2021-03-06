package models

import (
	"testing"
)

func NewReview(stars int, comment string) *CreateReviewCmd {

	return &CreateReviewCmd{stars, comment}
}

func Test_withCorrectParams(t *testing.T) {
	r := NewReview(4, "The iphone x looks good")
	err := r.validate()
	if err != nil {
		t.Error("the validation did not pass")
		t.Fail()
	}
}
func Test_shouldFailWithWrongNumberOfStars(t *testing.T) {
	r := NewReview(8, "The iphone looks REALLY good")

	err := r.validate()

	if err == nil {
		t.Error("should fail with 5 stars")
		t.Fail()
	}
}
