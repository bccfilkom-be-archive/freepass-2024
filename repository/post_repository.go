package repository

import (
	"github.com/nathakusuma/bcc-be-freepass-2024/entity"
	"github.com/nathakusuma/bcc-be-freepass-2024/model"
	"gorm.io/gorm"
	"time"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

func (repo *PostRepository) FindById(id uint) (*entity.Post, error) {
	var post entity.Post
	if err := repo.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (repo *PostRepository) FindByCandidateId(candId uint) ([]entity.Post, error) {
	var posts []entity.Post
	if err := repo.db.Where("candidate_id = ?", candId).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (repo *PostRepository) Create(post *entity.Post) error {
	return repo.db.Create(&post).Error
}

func (repo *PostRepository) Update(post *entity.Post, updates *model.UpdatePostRequest) error {
	return repo.db.Model(post).Updates(updates).UpdateColumn("updated_at", time.Now()).Error
}
