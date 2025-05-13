package auth

import (
	"errors"
	"example/crud/models"

	"golang.org/x/crypto/bcrypt"
)

// Service handles authentication business logic
type Service struct {
	// In-memory user storage (temporary, will be replaced with database)
	users map[string]*models.User
}

// NewService creates a new auth service
func NewService() *Service {
	return &Service{
		users: make(map[string]*models.User),
	}
}

// Register creates a new user
func (s *Service) Register(req *models.RegisterRequest) (*models.User, error) {
	// Check if user already exists
	if _, exists := s.users[req.Email]; exists {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create new user
	user := &models.User{
		ID:       uint(len(s.users) + 1), // Temporary ID generation
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Store user
	s.users[req.Email] = user

	return user, nil
}

// Login authenticates a user
func (s *Service) Login(req *models.LoginRequest) (*models.User, error) {
	// Find user
	user, exists := s.users[req.Email]
	if !exists {
		return nil, errors.New("invalid credentials")
	}

	// Check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
