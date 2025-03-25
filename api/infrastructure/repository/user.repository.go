package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"shanepee.com/api/domain"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func (a *userRepositoryImpl) CreateUser(ctx context.Context, user *domain.User) error {
	err := a.db.Create(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.ErrUserEmailAlreadyExist
		}
		return err
	}
	return err
}

func (u *userRepositoryImpl) UpdateUser(ctx context.Context, id int64, user map[string]any) error {
	if err := u.db.Model(domain.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepositoryImpl) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := u.db.Where("email = ?", email).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepositoryImpl) FindUserByID(ctx context.Context, id int64) (*domain.User, error) {
	var user domain.User
	if err := u.db.Take(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepositoryImpl) FindSellers(ctx context.Context) ([]*domain.UserWithReview, error) {
	var users []*domain.UserWithReview
	err := u.db.Model(&domain.User{}).
		Select("*, users.id AS id, users.photo AS photo, COALESCE(AVG(reviews.rating), 0) AS rating, COUNT(reviews.id) AS number_of_reviews").
		Joins("LEFT JOIN orders ON orders.seller_id = users.id").
		Joins("LEFT JOIN reviews ON reviews.order_id = orders.id").
		Group("users.id").
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepositoryImpl) FindSellerByID(ctx context.Context, id int64) (*domain.UserWithReview, error) {
	var user *domain.UserWithReview
	err := u.db.Model(&domain.User{}).
		Select("*, users.id AS id, users.photo AS photo, COALESCE(AVG(reviews.rating), 0) AS rating, COUNT(DISTINCT reviews.id) AS number_of_reviews, COUNT(order_items.id) AS number_of_art_toys_sold").
		Joins("LEFT JOIN orders ON orders.seller_id = users.id").
		Joins("LEFT JOIN order_items ON order_items.order_id = orders.id").
		Joins("LEFT JOIN reviews ON reviews.order_id = orders.id").
		Group("users.id").
		Take(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	var totalArtToys int64
	err = u.db.Model(&domain.ArtToy{}).
		Where("art_toys.owner_id = ?", id).
		Count(&totalArtToys).Error
	if err != nil {
		return nil, err
	}

	user.TotalArtToys = int(totalArtToys)
	return user, nil
}

func (u *userRepositoryImpl) CreatePasswordResetRequest(ctx context.Context, passwordResetRequest *domain.PasswordResetRequest) error {
	return u.db.Create(passwordResetRequest).Error
}

func (u *userRepositoryImpl) FindPasswordResetRequestWithUserByID(ctx context.Context, id int64) (*domain.PasswordResetRequest, error) {
	var passwordResetRequest domain.PasswordResetRequest
	if err := u.db.Joins("User").Take(&passwordResetRequest, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrPasswordResetRequestNotFound
		}
		return nil, err
	}
	return &passwordResetRequest, nil
}

func (u *userRepositoryImpl) UpdateUserPasswordHash(ctx context.Context, id int64, passwordHash string) error {
	user := &domain.User{ID: id}
	return u.db.Model(user).Updates(domain.User{PasswordHash: passwordHash}).Error
}

func (u *userRepositoryImpl) DeletePasswordResetRequestByID(ctx context.Context, id int64) error {
	return u.db.Delete(&domain.PasswordResetRequest{}, id).Error
}

var _ domain.UserRepository = &userRepositoryImpl{}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepositoryImpl{
		db,
	}
}
