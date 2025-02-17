package service

import (
	"context"
	"encoding/base64"
	"errors"

	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
	"shanepee.com/api/domain"
)

type AuthService interface {
	Register(ctx context.Context, username string, password string) (*domain.User, error)
	Login(ctx context.Context, email string, password string) (*domain.User, error)
	RequestPasswordChange(ctx context.Context, email string) error
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
	ChangePassword(ctx context.Context, requestID int64, token string, newPassword string) error
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

var (
	ErrIncorrectCredential error = errors.New("Invalid email or password")
	ErrInvalidToken        error = errors.New("Invalid token or request id")
)

func (s *authServiceImpl) Register(ctx context.Context, username string, password string) (*domain.User, error) {
	passwordByte := []byte(password)
	// TODO: salt
	// TODO: validate password
	hash, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	hashStr := string(hash)
	if err != nil {
		// TODO: properly handle this error
		return nil, err
	}
	user := domain.NewUser(username, hashStr)
	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		// TODO: properly handle this error
		return nil, err
	}

	return user, nil
}

func (s *authServiceImpl) Login(ctx context.Context, email string, password string) (*domain.User, error) {
	user, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, ErrIncorrectCredential
		}
		return nil, err
	}

	passwordByte := []byte(password)
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), passwordByte)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, ErrIncorrectCredential
		}
		return nil, err
	}
	return user, nil
}

func (s *authServiceImpl) RequestPasswordChange(ctx context.Context, email string) error {
	user, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			// This avoids exposing whether an email/user exists in the system.
			return nil
		}
		return err
	}

	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return err
	}
	token := base64.URLEncoding.EncodeToString(randomBytes)

	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	tokenHash := string(hash)
	passwordChangeRequest := domain.NewPasswordChangeRequest(tokenHash, user.ID)
	if err := s.userRepo.CreatePasswordChangeRequest(ctx, passwordChangeRequest); err != nil {
		return err
	}

	if err := s.emailSender.SendChangePasswordEmail(ctx, user.Email, token, passwordChangeRequest.ID); err != nil {
		return err
	}
	return nil
}

func (s *authServiceImpl) ChangePassword(ctx context.Context, requestID int64, token string, newPassword string) error {
	passwordChangeRequest, err := s.userRepo.FindPasswordChangeRequestWithUserByID(ctx, requestID)
	if err != nil {
		if errors.Is(err, domain.ErrPasswordChangeRequestNotFound) {
			return ErrInvalidToken
		}
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordChangeRequest.TokenHash), []byte(token)); err != nil {
		if errors.Is(err, domain.ErrPasswordChangeRequestNotFound) {
			return ErrInvalidToken
		}
		return err
	}

	passwordByte := []byte(newPassword)
	hash, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	hashStr := string(hash)
	if err != nil {
		return err
	}

	if err := s.userRepo.UpdateUserPasswordHash(ctx, passwordChangeRequest.UserID, hashStr); err != nil {
		return err
	}

	if err := s.userRepo.DeletePasswordChangeRequestByID(ctx, passwordChangeRequest.ID); err != nil {
		return err
	}
	return nil
}

func (s *authServiceImpl) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	user, err := s.userRepo.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
