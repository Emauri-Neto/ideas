package pg

func createThreadTable() string {
	return `
		CREATE TABLE thread(
			id UUID PRIMARY KEY,
			name TEXT NOT NULL,
			deadline TIMESTAMPTZ NOT NULL,
			study_id UUID REFERENCES study(id) ON DELETE CASCADE,
			responsible_user UUID REFERENCES users(id) ON DELETE CASCADE,
			created_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP - INTERVAL '3 hours'),
    		updated_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP - INTERVAL '3 hours')
		)
	`
}

func createThreadUsersTable() string {
	return `
		CREATE TABLE thread_users (
    	thread_id UUID REFERENCES thread(id) ON DELETE CASCADE,
   		user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    	PRIMARY KEY (thread_id, user_id)
		)
	`
}

func (d Driver) CreateThread() string {
	return "INSERT INTO thread(id, name, deadline, study_id, responsible_user) VALUES ($1, $2, $3, $4, $5)"
}