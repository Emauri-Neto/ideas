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

func (d Driver) CreateUser() string {
	return `
		INSERT INTO users(id, email, password, name) VALUES ($1, $2, $3, $4)
	`
}

func (d Driver) CreateStudy() string {
	return `
		INSERT INTO study(id, name, responsible_id) VALUES ($1, $2, $3)
	`
}

func (d Driver) IsStudyOwner() string {
	return `
		SELECT 1 FROM study WHERE id = $1 AND responsible_id = $2
	`
}

func (d Driver) CreateThread() string {
	return `
		INSERT INTO discussion_thread(id, name, study_id, responsible_id) VALUES ($1, $2, $3, $4)
	`
}
