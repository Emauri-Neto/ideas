package db

import "server/types"

func (db *Database) CreateStudy(s types.Study) (types.Study, error){
	var st types.Study
	err := db.sqlx.QueryRowx(db.query.CreateStudy(), s.Id, s.Title, s.Objective, s.Methodology, s.MaxParticipants, s.Private, s.UserID).StructScan(&st)

	if err != nil {
		return types.Study{}, err
	}

	return st, nil
}

func (db *Database) ListStudies(id string) ([]types.Study, error) {
	st := []types.Study{}

	if err := db.sqlx.Select(&st, db.query.ListStudies(), id); err != nil {
		return st, err
	}

	return st, nil
}
func (db *Database) GetStudy(id string) (types.Study, error){
	st := types.Study{}

	if err := db.sqlx.Get(&st, db.query.GetStudy(), id); err != nil {
		return st, err
	}

	return st, nil
}