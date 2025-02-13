package service

import (
	"context"
	"encoding/base64"
	"errors"

	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
	"shanepee.com/api/apperror"
	"shanepee.com/api/domain"
)

type AuthService interface {
	Register(ctx context.Context, username string, password string) (*domain.User, apperror.AppError)
	Login(ctx context.Context, email string, password string) (*domain.User, apperror.AppError)
	RequestPasswordChange(ctx context.Context, email string) apperror.AppError
	GetUserByID(ctx context.Context, id int64) (*domain.User, apperror.AppError)
	ChangePassword(ctx context.Context, requestID int64, token string, newPassword string) apperror.AppError
}

type EmailSender interface {
	SendChangePasswordEmail(ctx context.Context, to string, token string, requestID int64) error
}

func NewAuthService(userRepo domain.UserRepository, emailSender EmailSender) AuthService {
	return &authServiceImpl{
		userRepo,
		emailSender,
	}
}

type authServiceImpl struct {
	userRepo    domain.UserRepository
	emailSender EmailSender
}

var _ AuthService = &authServiceImpl{}

func (s *authServiceImpl) Register(ctx context.Context, username string, password string) (*domain.User, apperror.AppError) {
	passwordByte := []byte(password)
	// TODO: salt
	// TODO: validate password
	hash, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	hashStr := string(hash)
	if err != nil {
		// TODO: properly handle this error
		return nil, apperror.ErrInternal(err)
	}
	user := domain.NewUser(username, hashStr)
	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		// TODO: properly handle this error
		return nil, apperror.ErrInternal(err)
	}

	return user, nil
}

func (s *authServiceImpl) Login(ctx context.Context, email string, password string) (*domain.User, apperror.AppError) {
	user, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, apperror.ErrUnauthorized("Invalid email or password")
		}
		return nil, apperror.ErrInternal(err)
	}

	passwordByte := []byte(password)
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), passwordByte)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, apperror.ErrUnauthorized("Invalid email or password")
		}
		return nil, apperror.ErrInternal(err)
	}
	return user, nil
}

func (s *authServiceImpl) RequestPasswordChange(ctx context.Context, email string) apperror.AppError {
	user, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			// This avoids exposing whether an email/user exists in the system.
			return nil
		}
		return apperror.ErrInternal(err)
	}

	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return apperror.ErrInternal(err)
	}
	token := base64.URLEncoding.EncodeToString(randomBytes)

	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return apperror.ErrInternal(err)
	}
	tokenHash := string(hash)
	passwordChangeRequest := domain.NewPasswordChangeRequest(tokenHash, user.ID)
	if err := s.userRepo.CreatePasswordChangeRequest(ctx, passwordChangeRequest); err != nil {
		return apperror.ErrInternal(err)
	}

	if err := s.emailSender.SendChangePasswordEmail(ctx, user.Email, token, passwordChangeRequest.ID); err != nil {
		return apperror.ErrInternal(err)
	}
	return nil
}

func (s *authServiceImpl) ChangePassword(ctx context.Context, requestID int64, token string, newPassword string) apperror.AppError {
	passwordChangeRequest, err := s.userRepo.FindPasswordChangeRequestWithUserByID(ctx, requestID)
	if err != nil {
		if errors.Is(err, domain.ErrPasswordChangeRequestNotFound) {
			return apperror.ErrUnauthorized("Invalid token or request id")
		}
		return apperror.ErrInternal(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordChangeRequest.TokenHash), []byte(token)); err != nil {
		if errors.Is(err, domain.ErrPasswordChangeRequestNotFound) {
			return apperror.ErrUnauthorized("Invalid token or request id")
		}
		return apperror.ErrInternal(err)
	}

	passwordByte := []byte(newPassword)
	hash, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	hashStr := string(hash)
	if err != nil {
		return apperror.ErrInternal(err)
	}

	if err := s.userRepo.UpdateUserPasswordHash(ctx, passwordChangeRequest.UserID, hashStr); err != nil {
		return apperror.ErrInternal(err)
	}

	if err := s.userRepo.DeletePasswordChangeRequestByID(ctx, passwordChangeRequest.ID); err != nil {
		return apperror.ErrInternal(err)
	}
	return nil
}

func (s *authServiceImpl) GetUserByID(ctx context.Context, id int64) (*domain.User, apperror.AppError) {
	user, err := s.userRepo.FindUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, apperror.ErrNotFound("user not found")
		}
		return nil, apperror.ErrInternal(err)
	}
	return user, nil
}
