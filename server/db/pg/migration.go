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