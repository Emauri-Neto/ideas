package pg

func createUserTable() string {
	return `
		CREATE TABLE users (
			id UUID PRIMARY KEY,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			_verified BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP - INTERVAL '3 hours'),
			updated_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP - INTERVAL '3 hours')
		);
	`
}

func (d Driver) GetUserByEmail() string {
	return "SELECT * FROM users WHERE email = $1"
}

func (d Driver) CreateUser() string {
	return "INSERT INTO users(id, email, password) VALUES ($1, $2, $3)"
}

func (d Driver) GetUserById() string {
	return "SELECT * FROM users WHERE id = $1"
}