package config

import db "github.com/0xatanda/goFeedNews/sql/database"

type APIConfig struct {
	DB *db.Queries
}
