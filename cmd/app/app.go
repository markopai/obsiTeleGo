package app

import (
	"log/slog"
	"obsiTeleGo/internal/botHandler"
	"obsiTeleGo/internal/logger"
	"obsiTeleGo/internal/repository"
	"os"
)

type database interface {
	Close() error
}

type App struct {
	Logger     *logger.Logger
	Log        *slog.Logger
	db         database
	Repo       repository.Repo
	BotHandler *botHandler.BotHandler
}

type Options struct {
	Repo string
}

func New(opt *Options) App {
	base := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	logger := initLogger(base)
	log := initAppLog(base)

	repo, db, err := initRepo(logger.Repo)

	if err != nil {

		log.Error("Init Repo Error", "error", err)
		panic(err)
	}

	botHandler := initBotHandler(logger.BotHandler, repo)

	return App{
		Logger:     logger,
		Log:        log,
		db:         db,
		Repo:       repo,
		BotHandler: botHandler,
	}
}

func initLogger(base *slog.Logger) *logger.Logger {
	return logger.New(base)
}

func initAppLog(base *slog.Logger) *slog.Logger {
	return base.With("logger", "app")
}

func (a *App) DBClose() error {
	if a.db != nil {
		return a.db.Close()
	}
	return nil
}

func initBotHandler(log *slog.Logger, repo repository.Repo) *botHandler.BotHandler {
	return botHandler.New(log, repo)
}
