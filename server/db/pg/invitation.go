package pg

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

func (d Driver) GetInvitationsByReceiver() string {
	return `
		SELECT DISTINCT i.id, i.text, i.type, i.study_id, i.thread_id, i.status
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
			WHERE id = $1 AND status = 'pending'
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
		SET status = $1
		WHERE id = $2
	`
}
