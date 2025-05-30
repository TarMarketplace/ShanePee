package service

import (
	"context"
	"encoding/base64"
	"errors"
	"regexp"

	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
	"shanepee.com/api/domain"
)

var (
	ErrUserEmailAlreadyExist error = domain.ErrUserEmailAlreadyExist
)

type AuthService interface {
	Register(ctx context.Context, email string, password string) (*domain.User, error)
	Login(ctx context.Context, email string, password string) (*domain.User, error)
	RequestPasswordReset(ctx context.Context, email string) error
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
	ResetPassword(ctx context.Context, requestID int64, token string, newPassword string) error
	ChangePassword(ctx context.Context, userID int64, oldPassword string, newPassword string) error
}

type EmailSender interface {
	SendResetPasswordEmail(ctx context.Context, to string, token string, requestID int64) error
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
	ErrUnauthorized        error = errors.New("Unauthorized access")
	ErrIncorrectCredential error = errors.New("Invalid email or password")
	ErrInvalidToken        error = errors.New("Invalid token or request id")
	ErrUserNotFound        error = domain.ErrUserNotFound
	ErrInvalidOldPassword  error = errors.New("Invalid old password")
	ErrInvalidEmail        error = errors.New("Invalid email format")
)

// isValidEmail checks if the provided email has a valid format
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func (s *authServiceImpl) Register(ctx context.Context, email string, password string) (*domain.User, error) {
	// Validate email format
	if !isValidEmail(email) {
		return nil, ErrInvalidEmail
	}

	passwordByte := []byte(password)
	// TODO: salt
	// TODO: validate password
	hash, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	hashStr := string(hash)
	if err != nil {
		// TODO: properly handle this error
		return nil, err
	}
	user := domain.NewUser(email, hashStr)
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

func (s *authServiceImpl) RequestPasswordReset(ctx context.Context, email string) error {
	user, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			// This avoids exposing whether an email/user exists in the system.
			return nil
		}
		return err
	}

	token := make([]byte, 72)
	if _, err := rand.Read(token); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	tokenHash := string(hash)
	passwordResetRequest := domain.NewPasswordResetRequest(tokenHash, user.ID)
	if err := s.userRepo.CreatePasswordResetRequest(ctx, passwordResetRequest); err != nil {
		return err
	}

	encodedToken := base64.RawURLEncoding.EncodeToString(token)
	if err := s.emailSender.SendResetPasswordEmail(ctx, user.Email, encodedToken, passwordResetRequest.ID); err != nil {
		return err
	}
	return nil
}

func (s *authServiceImpl) ResetPassword(ctx context.Context, requestID int64, token string, newPassword string) error {
	passwordResetRequest, err := s.userRepo.FindPasswordResetRequestWithUserByID(ctx, requestID)
	if err != nil {
		if errors.Is(err, domain.ErrPasswordResetRequestNotFound) {
			return ErrInvalidToken
		}
		return err
	}

	decodedToken, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return ErrInvalidToken
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordResetRequest.TokenHash), decodedToken); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
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

	if err := s.userRepo.UpdateUserPasswordHash(ctx, passwordResetRequest.UserID, hashStr); err != nil {
		return err
	}

	if err := s.userRepo.DeletePasswordResetRequestByID(ctx, passwordResetRequest.ID); err != nil {
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

func (s *authServiceImpl) ChangePassword(ctx context.Context, userID int64, oldPassword string, newPassword string) error {
	user, err := s.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return ErrInvalidOldPassword
		}
		return err
	}

	passwordByte := []byte(newPassword)
	hash, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	hashStr := string(hash)
	if err != nil {
		return err
	}

	if err := s.userRepo.UpdateUserPasswordHash(ctx, userID, hashStr); err != nil {
		return err
	}

	return nil
}
