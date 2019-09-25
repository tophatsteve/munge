package reverse

import (
	"log"

	"golang.org/x/net/context"
)

type Service interface {
	Reverse(ctx context.Context, in *ReverseRequest) (*ReverseResponse, error)
}

type server struct {
}

func NewService() Service {
	return server{}
}

func (s server) Reverse(ctx context.Context, in *ReverseRequest) (*ReverseResponse, error) {
	resp := ReverseResponse{
		Text: s.reverse(in.Text),
	}
	return &resp, nil
}

func (s server) reverse(input string) string {
	log.Printf("Return called with %s", input)
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	var returnString = string(runes)
	log.Printf("Retunring %s", returnString)
	return returnString
}
