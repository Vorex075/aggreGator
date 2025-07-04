package commands

import (
	"database/sql"

	"github.com/Vorex075/aggreGator/internal/config"
	"github.com/Vorex075/aggreGator/internal/database"
)

type State struct {
	cfg *config.Config
	db  *database.Queries
}

func NewState(cfg *config.Config) *State {
	db, _ := sql.Open("postgres", cfg.DbUrl)
	return &State{
		cfg: cfg,
		db:  database.New(db),
	}
}

func (s *State) GetDB() *database.Queries {
	return s.db
}

func (s *State) GetCurrentUser() string {
	return s.cfg.CurrentUser
}
