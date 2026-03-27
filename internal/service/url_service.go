package service

import (
	"errors"
	"time"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
	"url-shortener/internal/utils"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) CreateShortURL(original, customAlias string, ttlMinutes int) (*model.URL, error) {
	var code string

	// Custom alias
	if customAlias != "" {
		_, err := s.repo.FindByCode(customAlias)
		if err == nil {
			return nil, errors.New("alias already taken")
		}
		code = customAlias
	} else {
		code = utils.GenerateShortCode(6)
	}

	// Expiration
	var expiredAt *time.Time
	if ttlMinutes > 0 {
		t := time.Now().Add(time.Duration(ttlMinutes) * time.Minute)
		expiredAt = &t
	}

	url := &model.URL{
		OriginalURL: original,
		ShortCode:   code,
		Clicks:      0,
		ExpiredAt:   expiredAt,
	}

	err := s.repo.Save(url)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *URLService) GetOriginalURL(code, userAgent, ip string) (*model.URL, error) {
	url, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}

	if url.ExpiredAt != nil && time.Now().After(*url.ExpiredAt) {
		return nil, errors.New("link expired")
	}

	go s.repo.IncrementClicks(code)
	go s.repo.LogClick(code, userAgent, ip)

	return url, nil
}

func (s *URLService) GetStats(code string) (int, error) {
	// cek dulu apakah URL ada
	_, err := s.repo.FindByCode(code)
	if err != nil {
		return 0, err
	}

	return s.repo.GetStats(code)
}
