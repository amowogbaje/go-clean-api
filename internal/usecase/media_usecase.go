package usecase

import (
	"go_clean_api/internal/domain"
	"strings"
)

type mediaUsecase struct {
	mediaRepo domain.MediaRepository
}

func NewMediaUsecase(repo domain.MediaRepository) domain.MediaUsecase {
	return &mediaUsecase{mediaRepo: repo}
}

func (u *mediaUsecase) ProcessUpload(filename string) (*domain.Media, error) {
	// Simple type detection based on extension
	ext := strings.ToLower(filename[strings.LastIndex(filename, ".")+1:])
	mediaType := "image"
	if ext == "mp4" || ext == "avi" || ext == "mov" || ext == "mkv" {
		mediaType = "video"
	}

	media := &domain.Media{
		URL:  "/uploads/" + filename,
		Type: mediaType,
	}

	err := u.mediaRepo.Create(media)
	if err != nil {
		return nil, err
	}
	return media, nil
}

func (u *mediaUsecase) GetAllMedia() ([]domain.Media, error) {
	return u.mediaRepo.FetchAll()
}