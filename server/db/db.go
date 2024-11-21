package db

import (
	"errors"
	"fmt"
	"server/db/pg"
	"server/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	sqlx *sqlx.DB
	query queries
}

type queries interface {
	Schema() []string
	GetIndexLastMigration() string
	SaveMigration() string

	CreateUser() string
	GetUserByEmail() string
	GetUserById() string

	CreateVerificationCode() string

	CreateSession() string
	GetSessionByID() string
	UpdateSession() string
	DeleteSessionById() string

	CreateStudy() string
	ListStudies() string
	GetStudy() string

	CreateThread() string
}

func MountDatabase() (*Database, error) {
	name, _name := utils.GetEnv("DB_NAME")
	host, _host := utils.GetEnv("DB_HOST")
	port, _port := utils.GetEnv("DB_PORT")
	pw, _pw := utils.GetEnv("DB_PASS")
	user, _user := utils.GetEnv("DB_USER")

	if err := errors.Join(_name, _host, _port, _pw, _user); err != nil {
		return nil, fmt.Errorf("error mounting database -> %s", err)
	}

	db, _db := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, pw, host, port, name))

	if _db != nil {
		return nil, fmt.Errorf("error connecting to database -> %s", _db)
	}

	return &Database{
		sqlx: db,
		query: queries(pg.Driver{}),
	}, nil
}

func (db *Database) Migrate() error {
	last_migration := -1

	_ = db.sqlx.Get(&last_migration, db.query.GetIndexLastMigration())
	for i, query := range db.query.Schema() {
		if i <= last_migration {
			continue
		}
		_, _m := db.sqlx.Exec(query)

		if _m != nil {
			return fmt.Errorf("migration failed caused by -> %w", _m)
		}

		_, _l := db.sqlx.Exec(db.query.SaveMigration(), i, query)

		if _l != nil {
			return fmt.Errorf("failed trying to save migrations [%d position]\ncaused by -> %s", i, _l)
		}
	}

	return nil
}