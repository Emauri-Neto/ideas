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

func createStudyTable() string {
	return `
		CREATE TABLE  study (
			id UUID PRIMARY KEY,
			name TEXT NOT NULL,
			objective TEXT,
			methodology TEXT,
			max_participants INTEGER,
			participation_type TEXT,
			created_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP - INTERVAL '3 hours'),
			
			responsible_id UUID NOT NULL,
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
			created_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP - INTERVAL '3 hours'),

			responsible_id UUID NOT NULL,
			study_id UUID NOT NULL,
			CONSTRAINT fk_responsible FOREIGN KEY (responsible_id) REFERENCES users (id) ON DELETE CASCADE,
			CONSTRAINT fk_study FOREIGN KEY (study_id) REFERENCES study (id) ON DELETE CASCADE
		)
	`
}

func createInvitationTable() string {
	return `
		CREATE TABLE invitation (
			id UUID PRIMARY KEY,
			text TEXT,
			type TEXT,
			accept BOOLEAN,
			study_id UUID NOT NULL,
			thread_id UUID NOT NULL,
			created_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP - INTERVAL '3 hours'),

			CONSTRAINT fk_study FOREIGN KEY (study_id) REFERENCES study (id) ON DELETE CASCADE,
			CONSTRAINT fk_thread FOREIGN KEY (thread_id) REFERENCES discussion_thread (id) ON DELETE CASCADE
		)
	`
}

func createUsersInvitationTable() string {
	return `
		CREATE TABLE users_invitation (
			id UUID PRIMARY KEY,

			invitation_id UUID NOT NULL,
			sender_id UUID NOT NULL,
			receiver_id UUID NOT NULL,

			CONSTRAINT fk_sender FOREIGN KEY (sender_id) REFERENCES users (id) ON DELETE CASCADE,
			CONSTRAINT fk_receiver FOREIGN KEY (receiver_id) REFERENCES users (id) ON DELETE CASCADE,
			CONSTRAINT fk_invitation FOREIGN KEY (invitation_id) REFERENCES invitation (id) ON DELETE CASCADE
		)
	`
}

func createUsersStudyTable() string {
	return `
		CREATE TABLE users_study (
			id UUID PRIMARY KEY,

			user_id UUID NOT NULL,
			study_id UUID NOT NULL,

			CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
			CONSTRAINT fk_study FOREIGN KEY (study_id) REFERENCES study (id) ON DELETE CASCADE
		)
	`
}

func createUsersThreadTable() string {
	return `
		CREATE TABLE users_thread (
			id UUID PRIMARY KEY,
			role TEXT NOT NULL,

			user_id UUID NOT NULL,
			thread_id UUID NOT NULL,

			CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
			CONSTRAINT fk_thread FOREIGN KEY (thread_id) REFERENCES discussion_thread (id) ON DELETE CASCADE
		)	
	`
}

func (d Driver) CreateMiddleTableUser() string {
	return `
		WITH ins1 AS (
			INSERT INTO users_study(id, user_id, study_id)
			SELECT $1, $2, $3
			WHERE NOT EXISTS (
				SELECT 1 
				FROM users_study 
				WHERE user_id = $2 AND study_id = $3
			)
			RETURNING id
		)
		INSERT INTO users_thread(id, user_id, thread_id, role)
		SELECT $4, $2, $5, $6
		WHERE NOT EXISTS (
			SELECT 1
			FROM users_thread
			WHERE user_id = $2 AND thread_id = $5
		)
	`
}
