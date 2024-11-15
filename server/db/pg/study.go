package pg

func createStudyTable() string {
	return `
		CREATE TABLE study(
			id UUID PRIMARY KEY,
			title TEXT NOT NULL,
            objective TEXT,
			methodology TEXT NOT NULL,
            num_participants INT,
			max_participants INT,
            participation_type TEXT,
            _private BOOLEAN DEFAULT FALSE,
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			created_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP - INTERVAL '3 hours'),
			updated_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP - INTERVAL '3 hours')
		)
	`
}

func (d Driver) CreateStudy() string {
	return `INSERT INTO study(id, title, objective, methodology, max_participants, _private, user_id)
            VALUES ($1, $2, $3, $4, $5, $6, $7)
            RETURNING id, title, objective, methodology, max_participants, _private, user_id, created_at, updated_at;
    `
}

func (d Driver) ListStudies() string {
	return "SELECT * FROM study WHERE (_private = false) OR (_private = true AND user_id = $1)"
}

func (d Driver) GetStudy() string {
	return "SELECT * FROM study WHERE id=$1"
}