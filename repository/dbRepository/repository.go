package dbRepository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddUser(user model.User) (*model.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	// return "", errors.New("not implemented") // TODO: replace this
	return &user, nil
}

func (r *Repository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err // TODO: replace this
}

// Photo methods
func (r *Repository) AddPhoto(photo model.Photo) (uint, error) {
	return 0, errors.New("not implemented") // TODO: replace this
}

func (r *Repository) GetPhotoByID(photoID uint) (model.Photo, error) {
	return model.Photo{}, errors.New("not implemented") // TODO: replace this
}

func (r *Repository) DeletePhoto(photoID uint) error {
	return errors.New("not implemented") // TODO: replace this
}

func (r *Repository) GetAllPhotosByUser(username string) ([]model.Photo, error) {
	return nil, errors.New("not implemented") // TODO: replace this
}
