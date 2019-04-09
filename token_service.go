package main

type Authable interface {
	Encode(data interface{}) (string, error)
	Decode(token string) (interface{}, error)
}

type TokenService struct {
	repo Repository
}

func (s *TokenService) Encode(data interface{}) (string, error) {
	return "", nil
}

func (s *TokenService) Decode(token string) (interface{}, error) {
	return "", nil
}
