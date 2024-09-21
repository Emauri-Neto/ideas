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
			expertise TEXT,
			interests TEXT[],
			roles TEXT[],
			status TEXT,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		)
	`
}
