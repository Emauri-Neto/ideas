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
	return `
		SELECT 
			s.id AS study_id, 
			s.title, 
			s.objective, 
			s.methodology,
			s.user_id as study_owner,
			s.max_participants, 
			s.num_participants, 
			s.participation_type, 
			s._private, 
			s.created_at AS study_created_at, 
			s.updated_at AS study_updated_at,
			t.id AS thread_id, 
			t.name AS thread_name, 
			t.deadline AS thread_deadline, 
			t.responsible_user AS responsible_user_id, 
			u.id AS user_id, 
			u.name AS user_name, 
			u.email AS user_email
		FROM study s
		LEFT JOIN thread t ON t.study_id = s.id
		LEFT JOIN users u ON u.id = t.responsible_user
		WHERE (_private = false) OR (_private = true AND user_id = $1)
		ORDER BY s.id, t.id;
	`
}

func (d Driver) GetStudy() string {
	return `
		SELECT 
			s.id AS study_id, 
			s.title, 
			s.objective, 
			s.methodology,
			s.user_id as study_owner,
			s.max_participants, 
			s.num_participants, 
			s.participation_type, 
			s._private, 
			s.created_at AS study_created_at, 
			s.updated_at AS study_updated_at,
			t.id AS thread_id, 
			t.name AS thread_name, 
			t.deadline AS thread_deadline, 
			t.responsible_user AS responsible_user_id, 
			u.id AS user_id, 
			u.name AS user_name, 
			u.email AS user_email
		FROM study s
		LEFT JOIN thread t ON t.study_id = s.id
		LEFT JOIN users u ON u.id = t.responsible_user
		WHERE s.id = $1
		ORDER BY t.id;
	`
}