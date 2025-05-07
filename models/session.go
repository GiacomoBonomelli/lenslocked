package models

import (
	"database/sql"
	"fmt"

	"github.com/GiacomoBonomelli/lenslocked/rand"
)

type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new session.
	// When look up a session this will be left empty, as we only
	// store the hash of a session token in our database and
	// we cannot reverse it into a raw token.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	token, err := rand.SessionToken() // restituisce il token di 32 bytes, stringato
	if err != nil {
		return nil, fmt.Errorf("create:%w", err)
	}
	// TODO: hash the session token
	session := Session{
		UserID: userID,
		Token:  token,
		// Set the TokenHash
	}

	// TODO: store the session in our DB
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO:Implement SessionService.User
	return nil, nil
}
