//go:build sqlite

package app

import (
	"database/sql"
	"fmt"
	"log/slog"
	"obsiTeleGo/internal/repository"
	"obsiTeleGo/internal/repository/sqliteRepo"
	"os"

	_ "modernc.org/sqlite"
)

func initRepo(log *slog.Logger) (repository.Repo, database, error) {
	dbPath := os.Getenv("SQLITE_PATH")

	db, err := sql.Open("sqlite", dbPath)

	if err != nil {
		return nil, nil, fmt.Errorf("open conn to sqlite error: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, nil, fmt.Errorf("ping sqlite error: %w", err)
	}

	return sqliteRepo.New(db, log), db, nil
}
