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

func (a *authServiceImpl) Register(ctx context.Context, username string, password string) (*domain.User, apperror.AppError) {
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
	err = a.userRepo.CreateUser(ctx, user)
	if err != nil {
		// TODO: properly handle this error
		return nil, apperror.ErrInternal(err)
	}

	return user, nil
}

func (a *authServiceImpl) Login(ctx context.Context, email string, password string) (*domain.User, apperror.AppError) {
	user, err := a.userRepo.FindUserByEmail(ctx, email)
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

func (a *authServiceImpl) RequestPasswordChange(ctx context.Context, email string) apperror.AppError {
	user, err := a.userRepo.FindUserByEmail(ctx, email)
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
	if err := a.userRepo.CreatePasswordChangeRequest(ctx, passwordChangeRequest); err != nil {
		return apperror.ErrInternal(err)
	}

	if err := a.emailSender.SendChangePasswordEmail(ctx, user.Email, token, passwordChangeRequest.ID); err != nil {
		return apperror.ErrInternal(err)
	}
	return nil
}

func (a *authServiceImpl) ChangePassword(ctx context.Context, requestID int64, token string, newPassword string) apperror.AppError {
	passwordChangeRequest, err := a.userRepo.FindPasswordChangeRequestWithUserByID(ctx, requestID)
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

	if err := a.userRepo.UpdateUserPasswordHash(ctx, passwordChangeRequest.UserID, hashStr); err != nil {
		return apperror.ErrInternal(err)
	}

	if err := a.userRepo.DeletePasswordChangeRequestByID(ctx, passwordChangeRequest.ID); err != nil {
		return apperror.ErrInternal(err)
	}
	return nil
}
