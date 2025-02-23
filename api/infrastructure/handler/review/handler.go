package review

import (
	"shanepee.com/api/service"
)

type ReviewHandler struct {
	reviewSvc service.ReviewService
}

func NewHandler(reviewSvc service.ReviewService) ReviewHandler {
	return ReviewHandler{
		reviewSvc,
	}
}
