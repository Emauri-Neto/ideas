package pg

func (d Driver) CreateStudy() string {
	return `
		INSERT INTO study(id, name, objective, methodology, num_participants, max_participants, responsible_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
}

func (d Driver) GetAllStudy() string {
	return `
		SELECT * FROM study
	`
}

func (d Driver) GetStudyById() string {
	return `
		SELECT * FROM study WHERE id =  $1 
	`
}

func (d Driver) IsStudyOwner() string {
	return `
		SELECT 1 FROM study WHERE id = $1 AND responsible_id = $2
	`
}

func (d Driver) DeleteStudy() string {
	return `
		DELETE FROM study WHERE id = $1
	`
}

func (d Driver) UpdateStudy() string {
	return `
		UPDATE study SET
		name = $1,
		objective = $2,
		methodology = $3,
		max_participants = $4,
		participation_type = $5
		WHERE id = $6
	`
}

func (d Driver) GetUsersByStudy() string {
	return `
        SELECT u.id, u.name, u.email FROM users_study us
        JOIN users u
        ON us.user_id = u.id
        WHERE study_id =  $1;
    `
}

func (d Driver) GetResponsibleAndStudyId() string {
	return `
		WITH select_thread AS (
			SELECT *
			FROM discussion_thread
			WHERE id = $1
		)
		SELECT st.id as study_id, st.responsible_id as study_responsible, dt.responsible_id as thread_responsible
		FROM study st
		JOIN select_thread dt ON st.id = dt.study_id;
	`
}