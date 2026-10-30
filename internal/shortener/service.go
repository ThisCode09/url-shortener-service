package shortener

import (
	"crypto/rand"
	"errors"
	"log/slog"
	"math/big"
	"net/url"
	"url-shortener/internal/storage"
)

type Service struct {
	store *storage.JSONStore
}

func NewService(store *storage.JSONStore) *Service {
	return &Service{
		store: store,
	}
}


func (s *Service) GenerateCode() (string, error) {

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 8)
	for i := range 8 {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil

}

func (s *Service) Shorten(rawURL string) (string, error) {
	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		slog.Error("Not a valid url")
		return "", errors.New("Provided url isn't valid")
	}
	var code string
	for{
		code, err = s.GenerateCode()
		if err!=nil{
			slog.Error("An unexpected error occured")
			return "",errors.New("unexpected error")
		}
		_,exists:=s.store.Get(code)
		if !exists{
			break
		}

	}
	err=s.store.Set(code,rawURL)
	if err!=nil{
		return "",err
	}
	return code, nil
}

func (s *Service) Resolve(code string)(string,bool){
	return s.store.Get(code)
}