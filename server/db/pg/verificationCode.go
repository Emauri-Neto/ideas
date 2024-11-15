package pg

func createVerificationCodeTable() string {
	return `
		CREATE TABLE verification_code (
			id UUID PRIMARY KEY,
			type TEXT NOT NULL,
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			created_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP - INTERVAL '3 hours'),
			expires_at TIMESTAMPTZ
		);
	`
}

func (d Driver) CreateVerificationCode() string {
	return "INSERT INTO verification_code(id, type, user_id, expires_at) VALUES ($1, $2, $3, $4)"
}