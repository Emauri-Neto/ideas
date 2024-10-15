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
		createUsersStudyTable(),
		createUsersThreadTable(),
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
		INSERT INTO study(id, name, objective, methodology, responsible_id) VALUES ($1, $2, $3, $4, $5)
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

func (d Driver) CreateThread() string {
	return `
		INSERT INTO discussion_thread(id, name, study_id, responsible_id) VALUES ($1, $2, $3, $4)
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

func (d Driver) ExistInvitationAndUser() string {
	return `
		WITH select_user AS (
			SELECT id
			FROM users
			WHERE id = $1
		) SELECT CASE
			WHEN EXISTS(SELECT 1 FROM select_user)
			THEN (
				SELECT EXISTS(
					SELECT 1 
						FROM users_invitation ui
						INNER JOIN invitation i ON ui.invitation_id = i.id
						WHERE ui.receiver_id = (SELECT id FROM select_user)
						AND i.thread_id = $2
						LIMIT 1
				)
			)
			ELSE NULL
			END;
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

func (d Driver) GetInvitationsByReceiver() string {
	return `
		SELECT DISTINCT i.id, i.text, i.type, i.study_id, i.thread_id, i.accept
		FROM invitation i
		JOIN users_invitation ui 
		ON ui.invitation_id = i.id 
		WHERE ui.receiver_id = $1;
	`
}

func (d Driver) GetInvitationOwner() string {
	return `
		WITH  invitations AS (
			SELECT *
			FROM invitation
			WHERE id = $1 AND accept IS NULL
		) SELECT i.*
			FROM invitations i
			JOIN users_invitation ui
			ON i.id = ui.invitation_id
			WHERE ui.receiver_id = $2
	`
}

func (d Driver) AcceptRefuseInvitation() string {
	return `
		UPDATE invitation
		SET accept = $1
		WHERE id = $2
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
