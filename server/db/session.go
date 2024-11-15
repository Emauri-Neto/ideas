package db

import (
	"server/types"
)

func (db *Database) CreateSession(s types.Session) error {
	_, err := db.sqlx.Exec(db.query.CreateSession(), s.Id, s.UserAgent, s.UserId)

	return err
}

func (db *Database) GetSessionByID(id string) (types.Session, error) {
	s := types.Session{}

	if err := db.sqlx.Get(&s,db.query.GetSessionByID(), id); err != nil {
		return s, err
	}

	return s, nil
}

func (db *Database) UpdateSession(id string) error {
	_, err := db.sqlx.Exec(db.query.UpdateSession(), id)

	return err
}

func (db *Database) DeleteSessionById(id string) error {
	_, err := db.sqlx.Exec(db.query.DeleteSessionById(), id)

	return err
}