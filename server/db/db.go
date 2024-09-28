package db

import (
	"errors"
	"fmt"
	"ideas/db/pg"
	"ideas/types"
	"ideas/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type queries interface {
	Schema() []string
	SaveMigration() string
	GetIndexLastMigration() string
	GetUsers() string
	GetUserByEmail() string
	CreateUser() string
}

type Database struct {
	sqlx  *sqlx.DB
	query queries
}

func (db *Database) GetUsers() ([]types.User, error) {
	var users []types.User

	if err := db.sqlx.Select(&users, db.query.GetUsers()); err != nil {
		return users, nil
	}

	return users, nil
}

func (db *Database) GetUserByEmail(email string) (types.User, error) {
	var user types.User

	if err := db.sqlx.Get(&user, db.query.GetUserByEmail(), email); err != nil {
		return user, err
	}

	return user, nil
}

func (db *Database) CreateAccount(u types.User) error {
	_, err := db.sqlx.Exec(db.query.CreateUser(), u.Id, u.Email, u.Password, u.Name)

	return err
}

func MountDB() (*Database, error) {
	name, _name := utils.GetEnv("DB_NAME")
	host, _host := utils.GetEnv("DB_HOST")
	port, _port := utils.GetEnv("DB_PORT")
	pw, _pw := utils.GetEnv("DB_PASS")
	user, _user := utils.GetEnv("DB_USER")

	if _err := errors.Join(_name, _host, _port, _pw, _user); _err != nil {
		return nil, fmt.Errorf("erro montando o banco de dados -> %s", _err)
	}

	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, pw, host, port, name))

	if err != nil {
		return nil, fmt.Errorf("erro conectando ao banco de dados -> %s", err)
	}

	return &Database{
		query: queries(pg.Driver{}),
		sqlx:  db,
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
			return fmt.Errorf("falha realizando a migration -> %w", _m)
		}

		_, _l := db.sqlx.Exec(db.query.SaveMigration(), i, query)

		if _l != nil {
			return fmt.Errorf("falha inserindo a migration no banco de dados -> posicao: %d \nerror: %s", i, query)
		}
	}

	return nil
}
