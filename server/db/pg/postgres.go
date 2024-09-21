package pg

type Driver struct{}

func (d Driver) Schema() []string {
	return []string{
		createMigrationTable(),
		createUserTable(),
	}
}

func (d Driver) GetUsers() string {
	return `
		SELECT id, name FROM users
	`
}

func (d Driver) GetUserByEmail() string {
	return `
		SELECT id, name, email, password FROM users WHERE email = $1
	`
}

func (d Driver) CreateUser() string {
	return `
		INSERT INTO users(id, email, password, name) VALUES ($1, $2, $3, $4)
	`
}
