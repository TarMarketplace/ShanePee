package arttoy

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"shanepee.com/api/domain"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/service"
)

type ArtToySortKey string

const (
	ArtToyReleaseDateSortKey = "release_date"
	ArtToyPriceSortKey       = "price"
)

type SearchArtToysInput struct {
	Keyword string        `query:"keyword"`
	SortKey ArtToySortKey `query:"sort_key" enum:"release_date,price" doc:"Sorting key. Available values: 'release_date', 'price'."`
	Reverse bool          `query:"reverse" doc:"If true, sorting is in descending order. Sorting is applied only if 'sort_key' is defined."`
}

type SearchArtToysOutput struct {
	Body handler.ArrayResponse[domain.ArtToy]
}

func (h *ArtToyHandler) RegisterSearchArtToys(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "search-art-toys",
		Method:      http.MethodGet,
		Path:        "/v1/art-toy/search",
		Tags:        []string{"Art toy"},
		Summary:     "Search Art Toys",
		Description: "Search art toys",
	}, func(ctx context.Context, i *SearchArtToysInput) (*SearchArtToysOutput, error) {
		var sortKey *service.ArtToySortKey

		switch i.SortKey {
		case ArtToyReleaseDateSortKey:
			value := service.ArtToyReleaseDateSortKey
			sortKey = &value
		case ArtToyPriceSortKey:
			value := service.ArtToyPriceSortKey
			sortKey = &value
		}

		searchParams := &service.ArtToySearchParams{
			Keyword: i.Keyword,
			SortKey: sortKey,
			Reverse: i.Reverse,
		}

		data, err := h.artToySvc.GetArtToysBySearchParams(ctx, searchParams)
		if err != nil {
			logrus.Error(err)
			return nil, handler.ErrInternalServerError
		}
		return &SearchArtToysOutput{
			Body: handler.ArrayResponse[domain.ArtToy]{Data: data},
		}, nil
	})
}
