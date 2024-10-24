package db

import (
	"database/sql"
	"errors"
	"fmt"
	"ideas/db/pg"
	"ideas/types"
	"ideas/utils"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type queries interface {
	Schema() []string
	SaveMigration() string
	GetIndexLastMigration() string
	GetUserByEmail() string
	CreateUser() string
	UpdateUser() string
	DeleteUser() string
	GetUserById() string
	CreateStudy() string
	GetAllStudy() string
	GetStudyById() string
	IsStudyOwner() string
	DeleteStudy() string
	UpdateStudy() string
	CreateThread() string

	GetUsersByStudy() string
	GetUsersByThread() string
	GetThreadById() string
	IsThreadResponsibleUnion() string
	ExistInvitationAndUser() string
	CreateInvitationWith() string
	GetResponsibleAndStudyId() string
	GetInvitationsByReceiver() string
	GetInvitationOwner() string
	AcceptRefuseInvitation() string
	CreateMiddleTableUser() string
}

type Database struct {
	sqlx  *sqlx.DB
	query queries
}

func (db *Database) GetUsersById(id string) (types.User, error) {
	var users types.User

	if err := db.sqlx.Get(&users, db.query.GetUserById(), id); err != nil {
		return users, err
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

func (db *Database) UpdateUser(u types.User) error {
	_, err := db.sqlx.Exec(db.query.UpdateUser(), u.Id, u.Name)

	return err
}

func (db *Database) DeleteUser(id string) error {
	_, err := db.sqlx.Exec(db.query.DeleteUser(), id)

	return err
}

func (db *Database) CreateStudy(s types.Study) error {
	_, err := db.sqlx.Exec(db.query.CreateStudy(), s.Id, s.Name, s.Objective, s.Methodology, s.Responsible_id)

	return err
}

func (db *Database) GetStudyById(id string) (types.Study, error) {
	var study types.Study

	if err := db.sqlx.Get(&study, db.query.GetStudyById(), id); err != nil {
		return study, err
	}
	return study, nil
}

func (db *Database) GetAllStudy() ([]types.Study, error) {
	studies := []types.Study{}

	if err := db.sqlx.Select(&studies, db.query.GetAllStudy()); err != nil {
		return studies, err
	}

	return studies, nil
}

func (db *Database) IsStudyOwner(id_study, id_user string) error {
	var exists int

	if err := db.sqlx.Get(&exists, db.query.IsStudyOwner(), id_study, id_user); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("usuário não é dono do estudo ou estudo não existe")
		}
		return err
	}

	if exists != 1 {
		return errors.New("falha ao verificar propriedade do estudo")
	}

	return nil
}

func (db *Database) DeleteStudy(id string) error {
	_, err := db.sqlx.Exec(db.query.DeleteStudy(), id)

	return err
}

func (db *Database) UpdateStudy(s types.Study) error {
	_, err := db.sqlx.Exec(db.query.UpdateStudy(), s.Name, s.Objective, s.Methodology, s.Max_participants, s.Participation_type, s.Id)

	return err
}

func (db *Database) CreateThread(t types.Thread) error {
	_, err := db.sqlx.Exec(db.query.CreateThread(), t.Id, t.Name, t.Study_id, t.Responsible_id)

	return err
}

func (db Database) GetUsersByStudy(study_id string) ([]types.UserResponse, error) {
	var users []types.UserResponse

	if err := db.sqlx.Select(&users, db.query.GetUsersByStudy(), study_id); err != nil {
		return users, err
	}

	if users == nil {
		return users, errors.New("nada encontrado")
	}

	return users, nil
}

func (db Database) GetUsersByThread(thread_id string) ([]types.UserResponse, error) {
	var users []types.UserResponse

	if err := db.sqlx.Select(&users, db.query.GetUsersByThread(), thread_id); err != nil {
		return users, err
	}

	if users == nil {
		return users, errors.New("nada encontrado")
	}

	return users, nil
}

func (db *Database) GetThreadById(id string) (types.Thread, error) {
	var thread types.Thread

	if err := db.sqlx.Get(&thread, db.query.GetThreadById(), id); err != nil {
		return thread, err
	}

	return thread, nil
}

func (db *Database) GetResponsibleAndStudyId(thread_id string) (types.Responsibles, error) {
	var responsibles types.Responsibles

	if err := db.sqlx.Get(&responsibles, db.query.GetResponsibleAndStudyId(), thread_id); err != nil {
		if err == sql.ErrNoRows {
			return responsibles, errors.New("não foi encontrado a thread")
		}
		return responsibles, err
	}

	return responsibles, nil
}

func (db *Database) ExistInvitationAndUser(thread_id, user_id string) error {
	var exist sql.NullBool

	if err := db.sqlx.Get(&exist, db.query.ExistInvitationAndUser(), user_id, thread_id); err != nil {
		return err
	}

	if exist.Valid {
		if exist.Bool {
			return errors.New("já existe convite para esse usuario")
		} else {
			return nil
		}
	}

	return errors.New("esse usuario não existe")
}

func (db *Database) GetInvitationsByReceiver(userReceiver string) ([]types.Invitation, error) {
	var invitations []types.Invitation

	if err := db.sqlx.Select(&invitations, db.query.GetInvitationsByReceiver(), userReceiver); err != nil {
		return nil, err
	}

	return invitations, nil
}

func (db *Database) CreateInvitation(invitation types.Invitation, userInvitation types.UserInvitation) error {
	_, err := db.sqlx.Exec(db.query.CreateInvitationWith(),
		invitation.Id, invitation.Type_invitation,
		invitation.Text, invitation.Study_id, invitation.Thread_id,
		userInvitation.Id, userInvitation.Sender_id, userInvitation.Receiver_id)

	return err
}

func (db *Database) GetInvitationOwner(invitation_id, user_id string) (types.Invitation, error) {
	var invitation types.Invitation

	if err := db.sqlx.Get(&invitation, db.query.GetInvitationOwner(), invitation_id, user_id); err != nil {
		if err == sql.ErrNoRows {
			return invitation, errors.New("esse usuario não é dono desse convite ou convite não existe")
		}
		return invitation, err
	}

	return invitation, nil
}

func (db *Database) AcceptRefuseInvite(invitation_id string, status string) error {
	_, err := db.sqlx.Exec(db.query.AcceptRefuseInvitation(), status, invitation_id)

	if err != nil {
		return err
	}

	return nil
}

func (db *Database) CreateMiddleTableUser(user_id string, invite types.Invitation) error {
	_, err := db.sqlx.Exec(db.query.CreateMiddleTableUser(),
		uuid.New().String(), user_id, invite.Study_id,
		uuid.New().String(), invite.Thread_id, invite.Type_invitation)

	if err != nil {
		return err
	}

	return nil
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
