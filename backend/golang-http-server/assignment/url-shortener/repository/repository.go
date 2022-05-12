package repository

import (
	"sync"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
)

type URLRepository struct {
	mu   sync.Mutex
	Data map[string]string
}

func NewMapRepository() URLRepository {
	data := make(map[string]string, 0)
	return URLRepository{
		Data: data,
	}
}

func (r *URLRepository) Get(path string) (*entity.URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if url, exist := r.Data[path]; exist {
		return &entity.URL{
			LongURL:  url,
			ShortURL: path,
		}, nil
	} else {
		return nil, entity.ErrURLNotFound
	}
}

func (r *URLRepository) Create(longURL string) (*entity.URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	randomShortURL := entity.GetRandomShortURL(longURL)
	r.Data[randomShortURL] = longURL

	return &entity.URL{
		LongURL:  longURL,
		ShortURL: randomShortURL,
	}, nil
}

func (r *URLRepository) CreateCustom(longURL, customPath string) (*entity.URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// checking
	if _, exist := r.Data[customPath]; exist {
		return nil, entity.ErrCustomURLIsExists
	} else {
		r.Data[customPath] = longURL

		return &entity.URL{
			LongURL:  longURL,
			ShortURL: customPath,
		}, nil
	}
}
