package pg

func (d Driver) CreateStudy() string {
	return `
		INSERT INTO study(id, name, responsible_id) VALUES ($1, $2, $3)
	`
}

func (d Driver) IsStudyOwner() string {
	return `
		SELECT 1 FROM study WHERE id = $1 AND responsible_id = $2
	`
}

func (d Driver) CreateThread() string {
	return `
		INSERT INTO discussion_thread(id, name, study_id, responsible_id) VALUES ($1, $2, $3, $4)
	`
}
