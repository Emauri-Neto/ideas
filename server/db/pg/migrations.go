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

func createTimeTable() string {
	return `
		CREATE TABLE time (
    		id UUID PRIMARY KEY,
    		date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    		week INTEGER,
    		month INTEGER,
    		year INTEGER
		)
	`
}

func createStudyTable() string {
	return `
		CREATE TABLE  study (
			id UUID PRIMARY KEY,
			name TEXT NOT NULL,
			objective TEXT,
			methodology TEXT,
			max_participants INTEGER,
			participation_type TEXT,
			time_id UUID NOT NULL,
			responsible_id UUID NOT NULL,
			CONSTRAINT fk_time FOREIGN KEY (time_id) REFERENCES time (id) ON DELETE CASCADE,
			CONSTRAINT fk_responsible FOREIGN KEY (responsible_id) REFERENCES users (id) ON DELETE CASCADE
		)
	`
}

func createDiscussionThreadTable() string {
	return `
		CREATE TABLE discussion_thread (
			id UUID PRIMARY KEY,
			name TEXT NOT NULL,
			max_participants INTEGER,
			discussion_deadline DATE,
			status TEXT,
			time_id UUID NOT NULL,
			responsible_id UUID NOT NULL,
			study_id UUID NOT NULL,
			CONSTRAINT fk_time FOREIGN KEY (time_id) REFERENCES time (id) ON DELETE CASCADE,
			CONSTRAINT fk_responsible FOREIGN KEY (responsible_id) REFERENCES users (id) ON DELETE CASCADE,
			CONSTRAINT fk_study FOREIGN KEY (study_id) REFERENCES study (id) ON DELETE CASCADE
		)
	`
}
