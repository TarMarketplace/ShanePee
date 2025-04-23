package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"shanepee.com/api/domain"
	"shanepee.com/api/service"
)

// MockUserRepository is a mock of domain.UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, id int64, user map[string]any) error {
	args := m.Called(ctx, id, user)
	return args.Error(0)
}

func (m *MockUserRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) FindUserByID(ctx context.Context, id int64) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) FindSellers(ctx context.Context) ([]*domain.UserWithReview, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.UserWithReview), args.Error(1)
}

func (m *MockUserRepository) FindSellerByID(ctx context.Context, id int64) (*domain.UserWithReview, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.UserWithReview), args.Error(1)
}

func (m *MockUserRepository) CreatePasswordResetRequest(ctx context.Context, passwordResetRequest *domain.PasswordResetRequest) error {
	args := m.Called(ctx, passwordResetRequest)
	return args.Error(0)
}

func (m *MockUserRepository) FindPasswordResetRequestWithUserByID(ctx context.Context, id int64) (*domain.PasswordResetRequest, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.PasswordResetRequest), args.Error(1)
}

func (m *MockUserRepository) UpdateUserPasswordHash(ctx context.Context, id int64, passwordHash string) error {
	args := m.Called(ctx, id, passwordHash)
	return args.Error(0)
}

func (m *MockUserRepository) DeletePasswordResetRequestByID(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockEmailSender is a mock of service.EmailSender
type MockEmailSender struct {
	mock.Mock
}

func (m *MockEmailSender) SendResetPasswordEmail(ctx context.Context, to string, token string, requestID int64) error {
	args := m.Called(ctx, to, token, requestID)
	return args.Error(0)
}

// TestCase structure to make tests DRY
type authRegisterTestCase struct {
	name      string
	email     string
	password  string
	wantErr   bool
	errType   error
	mockSetup func(*MockUserRepository, *MockEmailSender)
}

// Helper function to run a test case
func runAuthRegisterTest(t *testing.T, tc authRegisterTestCase) {
	// Setup mocks
	userRepo := new(MockUserRepository)
	emailSender := new(MockEmailSender)

	// Apply mock setup
	tc.mockSetup(userRepo, emailSender)

	// Create auth service with our custom implementation that includes email validation
	authService := service.NewAuthService(userRepo, emailSender)

	// Call Register
	user, err := authService.Register(context.Background(), tc.email, tc.password)

	// Assertions
	if tc.wantErr {
		assert.Error(t, err)
		if tc.errType != nil {
			assert.Equal(t, tc.errType.Error(), err.Error())
		}
		assert.Nil(t, user)
	} else {
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, tc.email, user.Email)
	}
}

// TC1-1: Successful registration
func TestAuthServiceImpl_Register_Success(t *testing.T) {
	tc := authRegisterTestCase{
		name:     "TC1-1: Successful registration",
		email:    "abcd@xyz.com",
		password: "12345678",
		wantErr:  false,
		mockSetup: func(userRepo *MockUserRepository, emailSender *MockEmailSender) {
			userRepo.On("CreateUser", mock.Anything, mock.MatchedBy(func(user *domain.User) bool {
				return user.Email == "abcd@xyz.com"
			})).Return(nil)
		},
	}

	runAuthRegisterTest(t, tc)
}

// TC1-2: Invalid email format (abcd)
func TestAuthServiceImpl_Register_InvalidEmailFormat(t *testing.T) {
	tc := authRegisterTestCase{
		name:     "TC1-2: Invalid email format (abcd)",
		email:    "abcd",
		password: "12345678",
		wantErr:  true,
		errType:  service.ErrInvalidEmail,
		mockSetup: func(userRepo *MockUserRepository, emailSender *MockEmailSender) {
			// No mock setup needed as validation should fail before repository is called
		},
	}

	runAuthRegisterTest(t, tc)
}

// TC1-3: Incorrect email format (@abcd.com)
func TestAuthServiceImpl_Register_IncorrectEmailFormat(t *testing.T) {
	tc := authRegisterTestCase{
		name:     "TC1-3: Incorrect email format (@abcd.com)",
		email:    "@abcd.com",
		password: "123456789",
		wantErr:  true,
		errType:  service.ErrInvalidEmail,
		mockSetup: func(userRepo *MockUserRepository, emailSender *MockEmailSender) {
			// No mock setup needed as validation should fail before repository is called
		},
	}

	runAuthRegisterTest(t, tc)
}

// TC1-4: Email already exists
func TestAuthServiceImpl_Register_EmailAlreadyExists(t *testing.T) {
	tc := authRegisterTestCase{
		name:     "TC1-4: Email already exists",
		email:    "exist@abcd.com",
		password: "123456789",
		wantErr:  true,
		errType:  domain.ErrUserEmailAlreadyExist,
		mockSetup: func(userRepo *MockUserRepository, emailSender *MockEmailSender) {
			userRepo.On("CreateUser", mock.Anything, mock.MatchedBy(func(user *domain.User) bool {
				return user.Email == "exist@abcd.com"
			})).Return(domain.ErrUserEmailAlreadyExist)
		},
	}

	runAuthRegisterTest(t, tc)
}

// bcrypt fail
func TestAuthServiceImpl_Register_BcryptFail(t *testing.T) {
	tc := authRegisterTestCase{
		name:     "TC1-4: Email already exists",
		email:    "exist@abcd.com",
		password: "12345678901234567890123456789012345678901234567890123456789012345678901234567890",
		wantErr:  true,
		mockSetup: func(userRepo *MockUserRepository, emailSender *MockEmailSender) {
			userRepo.On("CreateUser", mock.Anything, mock.MatchedBy(func(user *domain.User) bool {
				return user.Email == "exist@abcd.com"
			})).Return(domain.ErrUserEmailAlreadyExist)
		},
	}

	runAuthRegisterTest(t, tc)
}
