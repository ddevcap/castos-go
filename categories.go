package castos

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CategoriesService provides access to all Castos podcast category options.
type CategoriesService service

// Category represents a Castos podcast category option.
type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// GetAll returns al available podcast categories.
func (service *CategoriesService) GetAll() ([]*Category, error) {
	path := fmt.Sprintf("/get-categories")

	req, err := service.c.newRequest(http.MethodGet, path, url.Values{}, nil)
	if err != nil {
		return nil, err
	}

	categoriesList := map[string]map[string]string{}

	err = service.c.do(req, &categoriesList)
	if err != nil {
		return nil, err
	}

	if _, exists := categoriesList["categories"]; !exists {
		return nil, errors.New("no categories found in response data")
	}

	categories := make([]*Category, 0)

	for id, name := range categoriesList["categories"] {
		categoryId, _ := strconv.ParseInt(id, 10, 64)

		categories = append(categories, &Category{
			Id:   categoryId,
			Name: name,
		})
	}

	return categories, nil
}
