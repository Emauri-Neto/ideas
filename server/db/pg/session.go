package pg

func createSessionTable() string {
	return `
		CREATE TABLE session (
			id UUID PRIMARY KEY,
			user_agent TEXT,
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			created_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP - INTERVAL '3 hours'),
			expires_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP + INTERVAL '30 days')
		);
	`
}

func (d Driver) CreateSession() string {
	return "INSERT INTO session(id, user_agent, user_id) VALUES ($1, $2, $3)"
}

func (d Driver) GetSessionByID() string {
	return "SELECT * FROM session WHERE id = $1"
}

func (d Driver) UpdateSession() string {
	return "UPDATE session SET expires_at = (CURRENT_TIMESTAMP + INTERVAL '30 days') WHERE id = $1"
}

func (d Driver) DeleteSessionById() string {
	return "DELETE FROM session WHERE id = $1"
}