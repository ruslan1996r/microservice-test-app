package services

import "fmt"

type AuthService struct {
	AuthUrl      string
	FetchService IFetchService
}

func NewAuthService(fetchService IFetchService, authUrl string) *AuthService {
	return &AuthService{
		AuthUrl:      authUrl,
		FetchService: fetchService,
	}
}

func (s *AuthService) Login(path string) (string, error) {
	url := fmt.Sprintf("%s%s", s.AuthUrl, path)

	token, err := s.FetchService.Post(url, []byte{})

	if err != nil {
		return "", err
	}

	return token, nil
}
