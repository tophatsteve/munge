package capitalise

import (
	"log"
	"strings"

	"golang.org/x/net/context"
)

type Service interface {
	Capitalise(ctx context.Context, in *CapitaliseRequest) (*CapitaliseResponse, error)
}

type server struct {
}

func NewService() Service {
	return server{}
}

func (s server) Capitalise(ctx context.Context, in *CapitaliseRequest) (*CapitaliseResponse, error) {
	resp := CapitaliseResponse{
		Text: s.capitalise(in.Text),
	}
	return &resp, nil
}

func (s server) capitalise(input string) string {
	log.Printf("Capitalise called with %s", input)
	var returnString = strings.ToUpper(input)
	log.Printf("Retunring %s", returnString)
	return returnString
}
