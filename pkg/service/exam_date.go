package service

import (
	"github.com/ahmetk3436/pkg/model"
	post "github.com/ahmetk3436/pkg/repository/post"
)

// PostService ...
type PostService struct {
	PostRepository *post.Repository
}

// NewPostService ...
func NewPostService(p *post.Repository) PostService {
	return PostService{PostRepository: p}
}

// FindByID ...
func (p *PostService) FindByID(id uint) (*model.ExamDate, error) {
	return p.PostRepository.FindByID(id)
}

// Save ...
func (p *PostService) Save(examDate *model.ExamDate) (*model.ExamDate, error) {
	return p.PostRepository.Save(examDate)
}

// Migrate ...
func (p *PostService) Migrate() error {
	return p.PostRepository.Migrate()
}
