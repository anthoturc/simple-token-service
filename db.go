package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ApiToken struct {
	Id        uuid.UUID
	TokenHash string
	IsEnabled bool
	CreatedAt time.Time
	ExpiresAt time.Time
}

type DbService struct {
	DB *sql.DB
}

func (db *DbService) InsertApiToken(tokenHash string) (*ApiToken, error) {
	row := db.DB.QueryRow(`
		INSERT INTO api_tokens (token_hash)
		VALUES ($1)
		RETURNING id, is_enabled, created_at, expires_at;`, tokenHash)

	apiToken := &ApiToken{
		TokenHash: tokenHash,
	}
	err := row.Scan(&apiToken.Id, &apiToken.IsEnabled, &apiToken.CreatedAt, &apiToken.ExpiresAt)
	if err != nil {
		return nil, fmt.Errorf("could not insert token: %w", err)
	}

	return apiToken, nil
}

func (db *DbService) GetApiToken(tokenHash string) (*ApiToken, error) {
	row := db.DB.QueryRow(`
		SELECT id, is_enabled, created_at, expires_at
			FROM api_tokens
		WHERE token_hash = $1;`, tokenHash)
	apiToken := &ApiToken{
		TokenHash: tokenHash,
	}
	err := row.Scan(&apiToken.Id, &apiToken.IsEnabled, &apiToken.CreatedAt, &apiToken.ExpiresAt)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve api token: %w", err)
	}

	return apiToken, nil
}
