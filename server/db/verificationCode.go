package db

import (
	"server/types"
)

func (db *Database) CreateVerificationCode(v types.VerificationCode) error{
	_, err := db.sqlx.Exec(db.query.CreateVerificationCode(), v.Id, v.VCType, v.UserId, v.ExpiresAt)

	return err
}