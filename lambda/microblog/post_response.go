package microblog

import (
	"errors"
	"log"
)

type PostResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewPostResponse(data interface{}, err error) *PostResponse {
	if err == nil {
		return &PostResponse{Data: data}
	}

	log.Print(err)

	if errors.Is(err, ErrPostNotFound) {
		return &PostResponse{Message: ErrPostNotFound.Error()}
	}
	if errors.Is(err, ErrImageNotFound) {
		return &PostResponse{Message: ErrImageNotFound.Error()}
	}

	return &PostResponse{Message: "microblog: unexpected error"}
}
