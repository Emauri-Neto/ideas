package pg

func (d Driver) GetUserByEmail() string {
	return `
		SELECT * FROM users WHERE email = $1
	`
}

func (d Driver) GetUserById() string {
	return `
		SELECT id, name, email FROM users WHERE id = $1
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