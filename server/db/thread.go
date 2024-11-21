package db

import "server/types"

func (db *Database) CreateThread(t types.Thread) error {
	_, err := db.sqlx.Exec(db.query.CreateThread(), t.Id, t.Name, t.DeadLine, t.StudyID, t.ResponsibleID)

	return err
}
