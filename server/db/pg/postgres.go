package pg

type Driver struct{}

func (d Driver) Schema() []string {
	return []string{
		createMigrationTable(),
		createUserTable(),
		createWorkspace(),
		createStudyTable(),
		createDiscussionThreadTable(),
	}
}

func (d Driver) GetUsers() string {
	return `
		SELECT id, name FROM users
	`
}

func (d Driver) GetUserByEmail() string {
	return `
		SELECT * FROM users WHERE email = $1
	`
}

func (d Driver) GetUserById() string {
	return `
		SELECT id, name FROM users WHERE id = $1
	`
}

func (d Driver) CreateUser() string {
	return `
		INSERT INTO users(id, email, password, name) VALUES ($1, $2, $3, $4)
	`
}

func (d Driver) UpdateUser() string {
	return `
		UPDATE users SET name = $2 WHERE id = $1
	`
}

func (d Driver) DeleteUser() string {
	return `
		DELETE FROM users WHERE id = $1
	`
}

func (d Driver) CreateStudy() string {
	return `
		INSERT INTO study(id, name, responsible_id) VALUES ($1, $2, $3)
	`
}

func (d Driver) GetAllStudy() string {
	return `
		SELECT * FROM study
	`
}

func (d Driver) GetStudyById() string {
	return `
		SELECT * FROM study WHERE id =  $1 
	`
}

func (d Driver) IsStudyOwner() string {
	return `
		SELECT 1 FROM study WHERE id = $1 AND responsible_id = $2
	`
}

func (d Driver) DeleteStudy() string {
	return `
		DELETE FROM study WHERE id = $1
	`
}

func (d Driver) UpdateStudy() string {
	return `
		UPDATE study SET
		name = $1,
		objective = $2,
		methodology = $3,
		max_participants = $4,
		participation_type = $5
		WHERE id = $6
	`
}

func (d Driver) CreateThread() string {
	return `
		INSERT INTO discussion_thread(id, name, study_id, responsible_id) VALUES ($1, $2, $3, $4)
	`
}
