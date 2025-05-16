package auth

import (
	"database/sql"
	"errors"
	"example/crud/database"
	"example/crud/models"

	"golang.org/x/crypto/bcrypt"
)

// Service handles authentication business logic
type Service struct {
	db *sql.DB
}

// NewService creates a new auth service
func NewService() *Service {
	return &Service{
		db: database.DB,
	}
}

// Register creates a new user
func (s *Service) Register(req *models.RegisterRequest) (*models.User, error) {
	// Check if user already exists
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", req.Email).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create new user
	result, err := s.db.Exec(
		"INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		req.Username,
		req.Email,
		string(hashedPassword),
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       uint(id),
		Username: req.Username,
		Email:    req.Email,
	}

	return user, nil
}

// Login authenticates a user
func (s *Service) Login(req *models.LoginRequest) (*models.User, error) {
	var user models.User
	var hashedPassword string

	err := s.db.QueryRow(
		"SELECT id, username, email, password FROM users WHERE email = ?",
		req.Email,
	).Scan(&user.ID, &user.Username, &user.Email, &hashedPassword)

	if err == sql.ErrNoRows {
		return nil, errors.New("invalid credentials")
	}
	if err != nil {
		return nil, err
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
