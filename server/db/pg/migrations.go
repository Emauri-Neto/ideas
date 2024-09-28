package pg

func createMigrationTable() string {
	return `
		CREATE TABLE migrations (
			id INT PRIMARY KEY,
			query TEXT NOT NULL
		)
	`
}

func (d Driver) GetIndexLastMigration() string {
	return "SELECT MAX(id) FROM migrations"
}

func (d Driver) SaveMigration() string {
	return "INSERT INTO migrations(id, query) VALUES ($1, $2)"
}

func createUserTable() string {
	return `
		CREATE TABLE users(
			id UUID PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			created_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP - INTERVAL '3 hours'),
    		updated_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP - INTERVAL '3 hours')
		)
	`
}

func createWorkspace() string {
	return `
		CREATE TABLE workspace(
			id UUID PRIMARY KEY,
			subject TEXT NOT NULL,
			isPrivate BOOLEAN DEFAULT TRUE,
			content TEXT NOT NULL,
			user_id UUID NOT NULL,

    		CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
		)
	`
}