package services

import (
	"encoding/json"
	"fmt"

	"gateway/internal/types"
)

type ResourcesService struct {
	ResourcesUrl string
	FetchService IFetchService
}

func NewResourcesService(fetchService IFetchService, resourcesUrl string) *ResourcesService {
	return &ResourcesService{
		ResourcesUrl: resourcesUrl,
		FetchService: fetchService,
	}
}

func (s *ResourcesService) GetUsers(path string) (users *[]types.User, err error) {
	url := fmt.Sprintf("%s%s", s.ResourcesUrl, path)

	rawUsers, err := s.FetchService.Get(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(rawUsers), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *ResourcesService) GetBooks(path string) (books *[]types.Book, err error) {
	url := fmt.Sprintf("%s%s", s.ResourcesUrl, path)

	rawBooks, err := s.FetchService.Get(url)
	if err != nil {
		return nil, err
	}

	// var books []types.User

	err = json.Unmarshal([]byte(rawBooks), &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}
