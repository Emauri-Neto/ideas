package pg

type Driver struct{}

func (d Driver) Schema() []string {
	return []string{
		createMigrationTable(),
		createUserTable(),
		createStudyTable(),
		createDiscussionThreadTable(),
		createInvitationTable(),
		createUsersInvitationTable(),
		createUsersStudyTable(),
		createUsersThreadTable(),
	}
}