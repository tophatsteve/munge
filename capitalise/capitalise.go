package capitalise

import (
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
	return strings.ToUpper(input)
}
