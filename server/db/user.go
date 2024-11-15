package db

import "server/types"

func (db *Database) GetUserByEmail(email string) (types.User, error) {
	user := types.User{}

	if err := db.sqlx.Get(&user, db.query.GetUserByEmail(), email); err != nil {
		return user, err
	}

	return user, nil
}

func (db *Database) GetUserById(id string) (types.User, error){
	user := types.User{}

	if err := db.sqlx.Get(&user, db.query.GetUserById(), id); err != nil {
		return user, err
	}

	return user, nil
}

func (db *Database) CreateAccount(u types.User) error {
	_, err := db.sqlx.Exec(db.query.CreateUser(), u.Id, u.Email, u.Password)

	return err
}