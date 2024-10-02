package pg

type Driver struct{}

func (d Driver) Schema() []string {
	return []string{
		createMigrationTable(),
		createUserTable(),
		createWorkspace(),
		createStudyTable(),
		createDiscussionThreadTable(),
		createInvitationTable(),
		createUsersInvitationTable(),
	}
}

func (d Driver) GetUsers() string {
	return `
		SELECT id, name FROM users
	`
}

func (d Driver) GetUserByEmail() string {
	return `
		SELECT * FROM users WHERE email = $1
	`
}

func (d Driver) GetUserById() string {
	return `
		SELECT id, name FROM users WHERE id = $1
	`
}

func (d Driver) CreateUser() string {
	return `
		INSERT INTO users(id, email, password, name) VALUES ($1, $2, $3, $4)
	`
}

func (d Driver) UpdateUser() string {
	return `
		UPDATE users SET name = $2 WHERE id = $1
	`
}

func (d Driver) DeleteUser() string {
	return `
		DELETE FROM users WHERE id = $1
	`
}

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

func (d Driver) CreateInvitationWith() string {
	return `
		WITH ins_invitation AS (
    		INSERT INTO invitation(id, type, text, study_id, thread_id)
    		VALUES ($1, $2, $3, $4, $5)
    		RETURNING id
		)
		INSERT INTO users_invitation(id, invitation_id, sender_id, receiver_id)
		VALUES ($6, (SELECT id FROM ins_invitation), $7, $8);
	`
}

func (d Driver) ExistInvitation() string {
	return `
		SELECT 1 
		FROM users_invitation ui
		INNER JOIN invitation i ON ui.invitation_id = i.id
		WHERE ui.receiver_id = $1
		AND i.thread_id = $2
		LIMIT 1
	`
}
