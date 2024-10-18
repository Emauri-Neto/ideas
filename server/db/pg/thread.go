package pg

func (d Driver) CreateThread() string {
	return `
		INSERT INTO discussion_thread(id, name, study_id, responsible_id) VALUES ($1, $2, $3, $4)
	`
}

func (d Driver) GetUsersByThread() string {
	return `
        SELECT u.id, u.name, u.email FROM users_thread ut
        JOIN users u
        ON u.id = ut.user_id
        WHERE thread_id = $1
    `
}

func (d Driver) IsThreadResponsibleUnion() string {
	return `
		SELECT 1 FROM discussion_thread WHERE id = $1 AND responsible_id = $2
		UNION
		SELECT 1 FROM study WHERE id = $3 AND responsible_id = $2
	`
}

func (d Driver) GetThreadById() string {
	return `
		SELECT id, name, study_id FROM discussion_thread WHERE id = $1
	`
}