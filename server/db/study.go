package db

import (
	"fmt"
	"server/types"
)

type Result struct {
	StudyID           string  `db:"study_id"`
	Title             string  `db:"title"`
	Objective         string  `db:"objective"`
	Methodology       *string `db:"methodology"`
	OwnerID           string  `db:"study_owner"`
	MaxParticipants   *int    `db:"max_participants"`
	NumParticipants   *int    `db:"num_participants"`
	ParticipationType *string `db:"participation_type"`
	Private           bool    `db:"_private"`
	StudyCreatedAt    string  `db:"study_created_at"`
	StudyUpdatedAt    string  `db:"study_updated_at"`
	ThreadID          *string `db:"thread_id"`
	ThreadName        *string `db:"thread_name"`
	ThreadDeadline    *string `db:"thread_deadline"`
	ResponsibleUserID *string `db:"responsible_user_id"`
	UserID            *string `db:"user_id"`
	UserName          *string `db:"user_name"`
	UserEmail         *string `db:"user_email"`
}

func (db *Database) CreateStudy(s types.Study) (types.Study, error) {
	var st types.Study
	err := db.sqlx.QueryRowx(db.query.CreateStudy(), s.Id, s.Title, s.Objective, s.Methodology, s.MaxParticipants, s.Private, s.UserID).StructScan(&st)

	if err != nil {
		return types.Study{}, err
	}

	return st, nil
}

func (db *Database) ListStudies(id string) ([]types.Study, error) {

	var st []Result

	if err := db.sqlx.Select(&st, db.query.ListStudies(), id); err != nil {
		return []types.Study{}, err
	}

	stMap := make(map[string]*types.Study)

	for _, row := range st {
		if _, exists := stMap[row.StudyID]; !exists {
			stMap[row.StudyID] = &types.Study{
				Id:                row.StudyID,
				Title:             row.Title,
				Objective:         row.Objective,
				UserID:            row.OwnerID,
				Methodology:       row.Methodology,
				MaxParticipants:   row.MaxParticipants,
				NumParticipants:   row.NumParticipants,
				ParticipationType: row.ParticipationType,
				Private:           row.Private,
				CreatedAt:         row.StudyCreatedAt,
				UpdatedAt:         row.StudyUpdatedAt,
				Threads:           &[]types.Thread{},
			}
		}

		if row.ThreadID != nil {
			thread := types.Thread{
				Id:            *row.ThreadID,
				Name:          *row.ThreadName,
				DeadLine:      *row.ThreadDeadline,
				ResponsibleID: *row.ResponsibleUserID,
				StudyID:       row.StudyID,
			}

			if row.UserID != nil {
				thread.ResponsibleID = *row.ResponsibleUserID
			}

			*stMap[row.StudyID].Threads = append(*stMap[row.StudyID].Threads, thread)
		}
	}

	var studies []types.Study
	for _, study := range stMap {
		studies = append(studies, *study)
	}

	return studies, nil
}

func (db *Database) GetStudy(id string) (types.Study, error) {
	var st []Result

	if err := db.sqlx.Select(&st, db.query.GetStudy(), id); err != nil {
		fmt.Println(err)
		return types.Study{}, err
	}

	study := types.Study{
		Id:                st[0].StudyID,
		Title:             st[0].Title,
		Objective:         st[0].Objective,
		Methodology:       st[0].Methodology,
		MaxParticipants:   st[0].MaxParticipants,
		NumParticipants:   st[0].NumParticipants,
		ParticipationType: st[0].ParticipationType,
		Private:           st[0].Private,
		CreatedAt:         st[0].StudyCreatedAt,
		UpdatedAt:         st[0].StudyUpdatedAt,
		Threads:           &[]types.Thread{},
	}

	for _, row := range st {
		if row.ThreadID != nil {
			thread := types.Thread{
				Id:            *row.ThreadID,
				Name:          *row.ThreadName,
				DeadLine:      *row.ThreadDeadline,
				ResponsibleID: *row.ResponsibleUserID,
				StudyID:       row.StudyID,
			}

			if row.UserID != nil {
				thread.ResponsibleID = *row.UserID
			}

			*study.Threads = append(*study.Threads, thread)
		}
	}

	return study, nil
}
